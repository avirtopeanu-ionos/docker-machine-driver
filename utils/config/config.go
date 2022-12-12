package config

import (
	"errors"
	"gopkg.in/yaml.v3"
	"io/fs"
	"io/ioutil"
	//"fmt"
	"os"
	"path"
	"runtime"
)

const (
	EnvPath 	 = "IONOSCLOUD_CONFIG"
	FallbackPath = "/.config/docker-machine-driver-ionoscloud/ionoscloud.yaml"
)

func getFilePath() string {
	configFile := os.Getenv(EnvPath)
	if configFile != "" {
		return configFile
	}

	return path.Join(userHomeDir(), FallbackPath)
}

func userHomeDir() string {
	if runtime.GOOS != "windows" {
		return os.Getenv("HOME")
	}
	// Windows below
	if winUserProfile := os.Getenv("USERPROFILE"); winUserProfile != "" {
		return winUserProfile
	}
	return os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
}

type DriverCredentials struct {
	Token string `yaml:"token,omitempty"`
	Username string `yaml:"username,omitempty"`
	Password string `yaml:"password,omitempty"`
}

func Load() (*DriverCredentials, error) {
	config := new(DriverCredentials)
	if raw, err := ioutil.ReadFile(getFilePath()); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil, nil
		}
		return nil, err
	} else if formatErr := yaml.Unmarshal(raw, config); formatErr != nil {
		return nil, formatErr
	}
	return config, nil
}