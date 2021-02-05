package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	AccessToken           string
	Environment           Environment
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

func makePath(path string) error {
	dir := filepath.Dir(path)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Config) Load(profile string) error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(GetConfigFolder(os.Getenv("XDG_CONFIG_HOME")))
	err := viper.ReadInConfig()

	if err != nil {
		log.Error(err)
		return err
	}

	viper.Unmarshal(&c)

	err = viper.UnmarshalKey(profile, &c)
	if err != nil {
		log.Errorf("Error while loading config: %v", err)
		return err
	}

	return err
}

func (c *Config) GetAccessToken() (string, error) {
	if c.AccessToken != "" {
		return c.AccessToken, nil
	} else {
		if c.Environment == Production {
			return c.ProductionAccessToken, nil
		} else {
			return c.SandboxAccessToken, nil
		}
	}

	return "", errors.New("Your Access Token has not been setup. Use `square init` to set your Access Key")
}

func (c *Config) GetBaseURL() string {
	switch c.Environment {
	case Sandbox:
		if c.SandboxBaseUrl != "" {
			return c.SandboxBaseUrl
		}
		return "https://connect.squareupsandbox.com"
	case Production:
		if c.ProductionBaseUrl != "" {
			return c.ProductionBaseUrl
		}
		return "https://connect.squareup.com"
	default:
		return "https://connect.squareupsandbox.com"
	}
}
