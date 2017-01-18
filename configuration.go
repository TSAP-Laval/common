package common

import (
	"github.com/kelseyhightower/envconfig"
)

// Configuration représente la Configuration
// du package common
type Configuration struct {
	DatabaseDriver   string
	ConnectionString string
	SeedDataPath     string
}

// GetConfig récupère et valide la configuration
// du package
func GetConfig() (*Configuration, error) {
	var c Configuration

	err := envconfig.Process("TSAP", &c)

	if err != nil {
		return nil, err
	}

	return &c, err
}
