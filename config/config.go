package config

import (
	"encoding/json"
	"fmt"
	"os"
)

// Configs application configs
type Configs struct {
	DB            DatabaseConfig
	StopTimeoutMS int
	WebServer     WebServerConfig
	SwaggerUI     SwaggerConfig
}

// DatabaseConfig db connection configs
type DatabaseConfig struct {
	ConnectionString string
	Name             string
	LogMode          bool
}

// SwaggerConfig swagger configs
type SwaggerConfig struct {
	PageTitle   string
	Host        string
	Description string
	Schemes     []string
}

// WebServerConfig server configs
type WebServerConfig struct {
	Port int
	GIN  GINConfig
}

// GINConfig gin configs
type GINConfig struct {
	ReleaseMode bool
	UseLogger   bool
	UseRecovery bool
}

// Init function to read the config file
func Init(path string) (config *Configs, err error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("config file reading error, err : %v", err)
	}

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("config file unmarshalling error, err : %v", err)
	}

	return
}
