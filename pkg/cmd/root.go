package cmd

import (
	"fmt"
	"os"

	"github.com/nickrobinson/square-cli/pkg/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GoReleaser will update based on git tags
var version = "dev"

// Profile is the cli configuration for the user
var Profile config.Profile

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:           "square",
	SilenceUsage:  true,
	SilenceErrors: true,
	Annotations: map[string]string{
		"get": "api",
	},
	Version: version,
	Short:   "A CLI to help you develop your application with Square",
	Long: `The Square CLI gives you tools to make integrating your application
with Square easier. You can interact with all Square Connect APIs using this tool.

Before you use the CLI, you'll need to configure it:
$ square init`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(Profile.InitConfig)

	rootCmd.PersistentFlags().StringVar(&Profile.AccessToken, "access-token", "", "The access token to use for authentication")
	rootCmd.PersistentFlags().StringVar(&Profile.ConfigFile, "config", "", "config file (default is $HOME/.config/square/config.toml)")
	rootCmd.PersistentFlags().StringVar(&Profile.ProfileName, "profile", "default", "the profile name to read from for config")
	rootCmd.PersistentFlags().StringVar(&Profile.LogLevel, "log-level", "info", "log level (debug, info, warn, error)")

	viper.SetEnvPrefix("square")
	viper.AutomaticEnv()

	rootCmd.AddCommand(newGetCmd().reqs.Cmd)
	rootCmd.AddCommand(newDeleteCmd().reqs.Cmd)
	rootCmd.AddCommand(newPutCmd().reqs.Cmd)
	rootCmd.AddCommand(newPostCmd().reqs.Cmd)

	rootCmd.AddCommand(newInitCmd().Cmd)
}
