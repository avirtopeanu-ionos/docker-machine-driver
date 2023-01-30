package utils

import (
	"fmt"
	"github.com/ionos-cloud/docker-machine-driver/internal/pointer"
	"github.com/ionos-cloud/docker-machine-driver/pkg/sdk_utils"
	sdkgo "github.com/ionos-cloud/sdk-go/v6"
	"golang.org/x/exp/maps"
	"strconv"
	"time"
)

func (c *Client) GetNat(datacenterId, natId string) (*sdkgo.NatGateway, error) {
	nat, _, err := c.NATGatewaysApi.DatacentersNatgatewaysFindByNatGatewayId(c.ctx, datacenterId, natId).Execute()
	if err != nil {
		return nil, sdk_utils.ShortenOpenApiErr(err)
	}
	return &nat, nil
}

// NatRuleMaker 's objective is to easily create many NatGatewayRules with similar properties but different open ports
type NatRuleMaker struct {
	rules             []sdkgo.NatGatewayRule
	defaultProperties sdkgo.NatGatewayRuleProperties
}

func NewNRM(publicIp, srcSubnet, targetSubnet string) NatRuleMaker {
	return NatRuleMaker{
		rules: make([]sdkgo.NatGatewayRule, 0),
		defaultProperties: sdkgo.NatGatewayRuleProperties{
			Name: pointer.To("Docker Machine NAT Rule"),
			Type: pointer.To(sdkgo.NatGatewayRuleType("SNAT")),
			//Protocol:     pointer.To(sdkgo.NatGatewayRuleProtocol("ALL")),
			SourceSubnet: &srcSubnet,
			TargetSubnet: &targetSubnet,
			PublicIp:     &publicIp,
		},
	}
}

func (nrm NatRuleMaker) Make() *[]sdkgo.NatGatewayRule {
	return &nrm.rules
}

func (nrm NatRuleMaker) OpenPort(protocol string, port int32) NatRuleMaker {
	return nrm.OpenPorts(protocol, port, port)
}

func (nrm NatRuleMaker) OpenPorts(protocol string, start int32, end int32) NatRuleMaker {
	rule := sdkgo.NatGatewayRule{
		Properties: &nrm.defaultProperties,
	}
	rule.Properties.Protocol = (*sdkgo.NatGatewayRuleProtocol)(&protocol)
	rule.Properties.TargetPortRange = &sdkgo.TargetPortRange{Start: &start, End: &end}
	nrm.rules = append(nrm.rules, rule)
	return nrm
}

func (c *Client) CreateNat(datacenterId string, publicIps []string, lansToGateways map[string][]string, subnet string) (*sdkgo.NatGateway, error) {
	var lans []sdkgo.NatGatewayLanProperties
	fmt.Printf("CreateNat(publicIps = %+v, lansMap = %+v, subnet = %s)\n", publicIps, lansToGateways, subnet)

	err := c.createLansIfNotExist(datacenterId, maps.Keys(lansToGateways))
	if err != nil {
		return nil, err
	}
	time.Sleep(5 * time.Second)
	for lanId, gatewayIps := range lansToGateways {
		id, err := strconv.ParseInt(lanId, 10, 32)
		if err != nil {
			return nil, err
		}
		// Unpack the map into NatGatewayLanProperties objects. https://api.ionos.com/docs/cloud/v6/#tag/NAT-Gateways/operation/datacentersNatgatewaysPost
		var ptrGatewayIps *[]string = nil
		if len(gatewayIps) > 1 || gatewayIps[0] != "" {
			// We do this check so that we don't set the GatewayIps property if it's empty. If the property is empty, a gateway IP is generated by the API.
			ptrGatewayIps = &gatewayIps
		}
		lans = append(lans, sdkgo.NatGatewayLanProperties{Id: pointer.To(int32(id)), GatewayIps: ptrGatewayIps})
		fmt.Printf("Created a NatGatewayLanProperties obj with Id: %d, GatewayIps: %+v\n", id, gatewayIps)
	}

	nrm := NewNRM(publicIps[0], subnet, subnet)
	nrm.
		OpenPort("TCP", 22).            // SSH
		OpenPort("UDP", 53).            // DNS
		OpenPort("TCP", 80).            // HTTP
		OpenPort("TCP", 179).           // Calico BGP Port
		OpenPort("TCP", 443).           //
		OpenPort("TCP", 2376).          // Node driver Docker daemon TLS port
		OpenPort("UDP", 4789).          // Flannel VXLAN overlay networking on Windows cluster
		OpenPort("TCP", 6443).          // Rancher Webhook
		OpenPort("TCP", 6783).          // Weave Port
		OpenPort("TCP", 8443).          // Rancher webhook
		OpenPort("UDP", 8472).          // Canal/Flannel VXLAN overlay networking
		OpenPort("TCP", 9099).          // Canal/Flannel livenessProbe/readinessProbe
		OpenPort("TCP", 9100).          // Default port required by Monitoring to scrape metrics from Linux node-exporters
		OpenPort("TCP", 9443).          // Rancher webhook
		OpenPort("TCP", 9796).          // Default port required by Monitoring to scrape metrics from Windows node-exporters
		OpenPort("TCP", 10254).         // Ingress controller livenessProbe/readinessProbe
		OpenPort("TCP", 10256).         //
		OpenPorts("TCP", 2379, 2380).   // etcd
		OpenPorts("UDP", 6783, 6784).   // Weave Port (UDP)
		OpenPorts("TCP", 10250, 10252). // Metrics server communication with all nodes API
		OpenPorts("TCP", 30000, 32767). //
		OpenPorts("UDP", 30000, 32767). //
		OpenPort("ALL", 0)              // Outbound
	rules := nrm.Make()

	nat, resp, err := c.NATGatewaysApi.DatacentersNatgatewaysPost(c.ctx, datacenterId).NatGateway(
		sdkgo.NatGateway{
			Properties: &sdkgo.NatGatewayProperties{
				Name:      pointer.To("NAT Docker Machine"),
				PublicIps: &publicIps,
				Lans:      &lans,
			},
			Entities: &sdkgo.NatGatewayEntities{
				Rules:    &sdkgo.NatGatewayRules{Items: rules},
				Flowlogs: nil,
			},
		},
	).Execute()
	if err != nil {
		return nil, err
	}
	fmt.Printf("created nat: %+v\n", nat)

	err = c.waitTillProvisioned(resp.Header.Get("location"))
	return &nat, err
}

func (c *Client) createLansIfNotExist(datacenterId string, lanIds []string) error {
	for _, lanid := range lanIds {
		_, resp, err := c.LANsApi.DatacentersLansFindById(c.ctx, datacenterId, lanid).Execute()
		if resp.StatusCode == 404 {
			// Run this before err check, as 404s throws an err.
			fmt.Printf("Creating LAN %s for NAT\n", lanid)
			_, err := c.CreateLan(datacenterId, "Docker Machine LAN (NAT)", false)
			if err != nil {
				return err
			}
			continue // breakpoint
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) RemoveNat(datacenterId, natId string) error {
	_, err := c.NATGatewaysApi.DatacentersNatgatewaysDelete(c.ctx, datacenterId, natId).Execute()
	if err != nil {
		return sdk_utils.ShortenOpenApiErr(err)
	}
	return nil
}
