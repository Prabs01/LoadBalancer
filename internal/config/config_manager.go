package config

import (
	"fmt"
	"log"
	"os"
	"gopkg.in/yaml.v3"
)


type ConfigManager struct {
	configFilePath string
	configData     map[string]interface{}
}

func NewConfigManager(configFilePath string) *ConfigManager {
	return &ConfigManager{
		configFilePath: configFilePath,
		configData:     make(map[string]interface{}),
	}
}

type Timeout struct {
	Read int `yaml:"read"`
	Dial int `yaml:"dial"`
	Write int `yaml:"write"`
	Idle int `yaml:"idle"`
}

type Health struct {
	Passive struct {
		Fall     int `yaml:"fall"`
		Cooldown int `yaml:"cooldown"`
	} `yaml:"passive"`
}

type RoutesMatch struct {
	PathPrefix string `yaml:"path_prefix"`
}

type Routes struct {
	Match RoutesMatch `yaml:"match"`
	Pool  string      `yaml:"pool"`
}

type Pool struct {
	Addr string `yaml:"addr"`
	Weight int    `yaml:"weight,omitempty"`
}

type Config struct {
	Mode     int `yaml:"mode"`
	Listen   string `yaml:"listen"`
	Balancer string `yaml:"balancer"`

	Timeout Timeout `yaml:"timeout"`
	Health  Health  `yaml:"health"`

	Routes []Routes         `yaml:"routes"`
	Pools  map[string][]Pool `yaml:"pools"`
}

func (cm *ConfigManager) LoadConfig() (*Config, error) {
	config := &Config{}

	data, err := os.ReadFile(cm.configFilePath)
	if err != nil {
		log.Printf("Error reading config file: %v", err)
		return nil, fmt.Errorf("failed to read config file: %v", err)
	}

	log.Printf("Config file content: %s", string(data))

	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Printf("Error unmarshaling config: %v", err)
		return nil, fmt.Errorf("failed to unmarshal config: %v", err)
	}

	validationErr := validateConfig(config)
	if validationErr != nil {
		log.Printf("Config validation error: %v", validationErr)
		return nil, fmt.Errorf("config validation failed: %v", validationErr)
	}		

	log.Printf("Config loaded successfully: %+v", config)

	return config, nil
}

func validateConfig(config *Config) error {

	addrSet := make(map[string]struct{})
	for _, pools := range config.Pools {
		for _, pool := range pools {
			//all weights must be positive integers
			if pool.Weight <= 0 {
				log.Printf("Invalid weight %d for pool %s", pool.Weight, pool.Addr)
				return fmt.Errorf("invalid weight %d for pool %s", pool.Weight, pool.Addr)
			}
			//check for duplicate addresses
			if _, exists := addrSet[pool.Addr]; exists {
				log.Printf("Duplicate address %s found in pools", pool.Addr)
				return fmt.Errorf("duplicate address %s found in pools", pool.Addr)
			}
			addrSet[pool.Addr] = struct{}{}
		}
	}

	//every route must reference a valid pool
	for _, route := range config.Routes {
		if _, exists := config.Pools[route.Pool]; !exists {
			log.Printf("Route references non-existent pool %s", route.Pool)
			return fmt.Errorf("route references non-existent pool %s", route.Pool)
		}
	}

	//timeout values must be valid
	if config.Timeout.Read < 0 || config.Timeout.Dial < 0 || config.Timeout.Write < 0 || config.Timeout.Idle < 0 {
		log.Printf("Invalid timeout values found")
		return fmt.Errorf("timeout values must be non-negative")
	}


	return nil	

}