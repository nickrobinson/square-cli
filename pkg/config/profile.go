package config

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
)

// Profile handles all things related to managing the project specific configurations
type Profile struct {
	SandboxAccessToken    string `mapstructure:"sandbox_access_token"`
	ProductionAccessToken string `mapstructure:"production_access_token"`
	SandboxBaseUrl        string `mapstructure:"sandbox_base_url"`
	ProductionBaseUrl     string `mapstructure:"production_base_url"`
}

// GetConfigFolder retrieves the folder where the config file is stored
func GetConfigFolder(xdgPath string) string {
	configPath := xdgPath

	log.WithFields(log.Fields{
		"prefix": "config.Profile.GetConfigFolder",
		"path":   configPath,
	}).Debug("Using config file")

	if configPath == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configPath = filepath.Join(home, ".config")
	}

	return filepath.Join(configPath, "square")
}
