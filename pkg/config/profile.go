package config

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/nickrobinson/square-cli/pkg/validators"
	log "github.com/sirupsen/logrus"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

// Profile handles all things related to managing the project specific configurations
type Profile struct {
	ConfigFile  string
	LogLevel    string
	ProfileName string
	AccessToken string
}

// GetConfigFolder retrieves the folder where the config file is stored
func (p *Profile) GetConfigFolder(xdgPath string) string {
	configPath := xdgPath

	log.WithFields(log.Fields{
		"prefix": "profile.Profile.GetConfigFolder",
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

// InitConfig reads in config file and ENV variables if set.
func (p *Profile) InitConfig() {
	logFormatter := &prefixed.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC1123,
	}

	log.SetFormatter(logFormatter)

	// Set log level
	switch p.LogLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	default:
		log.Fatalf("Unrecognized log level value: %s. Expected one of debug, info, warn, error.", p.LogLevel)
	}

	if p.ConfigFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(p.ConfigFile)
	} else {
		configFolder := p.GetConfigFolder(os.Getenv("XDG_CONFIG_HOME"))
		configFile := filepath.Join(configFolder, "config.toml")
		viper.SetConfigType("toml")
		viper.SetConfigFile(configFile)
	}

	// If a profiles file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.WithFields(log.Fields{
			"prefix": "config.Profile.InitConfig",
			"path":   viper.ConfigFileUsed(),
		}).Debug("Using profiles file")
	}
}

func (p *Profile) GetAccessToken() (string, error) {
	key := viper.GetString("default.access_token")
	if key != "" {
		err := validators.AccessToken(key)
		if err != nil {
			return "", err
		}
		return key, nil
	}

	return "", errors.New("your Access Token has not been configured. Use `square init` to set your Access Key")
}
