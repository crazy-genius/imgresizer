package configuration

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// VERSION contain version of application
const VERSION = "development"

// Configuration represents configuration of application
type Configuration struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`

	EnableHTTPS        bool `yaml:"enabelHttps"`
	EnableRateLimit    bool `yaml:"enableRateLimits"`
	EnableImageStorage bool `yaml:"enableImageStorage"`
	MaxConnections     int  `yaml:"maxConnections"`

	AllowedHosts []string `yaml:"allowedHosts"`
}

// LoadConfiguration loads configuration data
func LoadConfiguration(configurationPath string) (*Configuration, error) {
	cfg := &Configuration{}

	f, err := os.Open(configurationPath)

	if err != nil {
		return nil, err
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
