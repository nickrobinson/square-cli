package config

import (
	"errors"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	ProfileName string `mapstructure:"profile"`
	ConfigFile  string
	LogLevel    string
	AccessToken string
	Environment Environment
	Profile     Profile
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

func (c *Config) Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	// v.AddConfigPath(GetConfigFolder(os.Getenv("XDG_CONFIG_HOME")))
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	if err != nil {
		log.Error(err)
		return err
	}

	viper.Unmarshal(&c)

	profile := viper.GetString("profile")
	err = viper.UnmarshalKey(profile, &c.Profile)
	if err != nil {
		log.Errorf("Error while loading config: %v", err)
		return err
	}

	log.Infof("Config: %v", c)

	return err
}

func (c *Config) GetProfile() *Profile {
	return &c.Profile
}

func (c *Config) GetAccessToken() (string, error) {
	if c.AccessToken != "" {
		return c.AccessToken, nil
	} else {
		if c.Environment == Production {
			return c.Profile.ProductionAccessToken, nil
		} else {
			return c.Profile.SandboxAccessToken, nil
		}
	}

	return "", errors.New("Your Access Token has not been setup. Use `square init` to set your Access Key")
}

func (c *Config) GetBaseURL() string {
	switch c.Environment {
	case Sandbox:
		if c.Profile.SandboxBaseUrl != "" {
			return c.Profile.SandboxBaseUrl
		}
		return "https://connect.squareupsandbox.com"
	case Production:
		if c.Profile.ProductionBaseUrl != "" {
			return c.Profile.ProductionBaseUrl
		}
		return "https://connect.squareupsandbox.com"
	default:
		return "https://connect.squareupsandbox.com"
	}
}
