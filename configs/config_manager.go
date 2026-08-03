package configs

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

	log.Printf("Config loaded successfully: %+v", config)

	return config, nil
}