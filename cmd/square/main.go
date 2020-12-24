package main

import (
	"github.com/nickrobinson/square-cli/internal/cmd"
	"github.com/nickrobinson/square-cli/pkg/square"
)

// GoReleaser will update based on git tags
var version = "dev"

func main() {
	cmd.Execute()
}

func buildRootCommand() *cobra.Cobra {
	s := square.New()
	cmd := &cobra.Command{
		Use:           "square",
		SilenceUsage:  true,
		SilenceErrors: true,
		Version:       version,
		Short:         "A CLI to help you develop your application with 🔲",
		Long: `The 🔲 CLI gives you tools to make integrating your application
	with Square easier. You can interact with all Square Connect APIs using this tool.
	
	Before you use the CLI, you'll need to configure it:
	$ square init`,
	}

	cmd.PersistentFlags().StringVar(&Profile.AccessToken, "access-token", "", "The access token to use for authentication")
	cmd.PersistentFlags().StringVar(&Profile.ConfigFile, "config", "", "config file (default is $HOME/.config/square/config.toml)")
	cmd.PersistentFlags().StringVarP(&Profile.ProfileName, "profile", "p", "default", "the profile name to read from for config")
	cmd.PersistentFlags().StringVar(&Profile.LogLevel, "log-level", "info", "log level (debug, info, warn, error)")
	cmd.PersistentFlags().VarP(&Profile.Environment, "env", "e", "Environment to use for request (sandbox/production)")

	return cmd
}
