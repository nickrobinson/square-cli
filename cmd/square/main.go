package main

import (
	"os"

	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// GoReleaser will update based on git tags
var version = "dev"

func main() {
	cmd := buildRootCommand()
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func buildRootCommand() *cobra.Command {
	s := square.New()

	cmd := &cobra.Command{
		Use:           "square",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			s.Config.Load()
			return nil
		},
		Short: "A CLI to help you develop your application with 🔲",
		Long: `The 🔲 CLI gives you tools to make integrating your application
		with Square easier. You can interact with all Square Connect APIs using this tool.
		
		Before you use the CLI, you'll need to configure it:
		$ square init`,
	}

	cmd.PersistentFlags().StringVar(&s.Config.AccessToken, "access-token", "", "The access token to use for authentication")
	// cmd.PersistentFlags().StringVar(&s.Config.ConfigFile, "config", "", "config file (default is $HOME/.config/square/config.toml)")
	cmd.PersistentFlags().StringP("profile", "p", "default", "the profile name to read from for config")
	cmd.PersistentFlags().StringVar(&s.Config.LogLevel, "log-level", "info", "log level (debug, info, warn, error)")
	cmd.PersistentFlags().VarP(&s.Config.Environment, "env", "e", "Environment to use for request (sandbox/production)")

	viper.SetDefault("environment", "sandbox")

	viper.BindPFlag("accessToken", cmd.PersistentFlags().Lookup("access-token"))
	viper.BindPFlag("profile", cmd.PersistentFlags().Lookup("profile"))
	viper.BindPFlag("logLevel", cmd.PersistentFlags().Lookup("log-level"))
	viper.BindPFlag("environment", cmd.PersistentFlags().Lookup("env"))

	cmd.AddCommand(buildGetCommand(s))
	cmd.AddCommand(buildPutCommand(s))
	cmd.AddCommand(buildPostCommand(s))
	cmd.AddCommand(buildDeleteCommand(s))

	cmd.AddCommand(buildStatusCommand(s))

	return cmd
}
