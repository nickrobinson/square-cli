package main

import (
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
)

func buildInitCommand(s *square.Square) *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Args:  cobra.NoArgs,
		Short: "Initialize Square CLI config.",
		Long:  `Initialize Square CLI configuration file.`,

		RunE: func(cmd *cobra.Command, args []string) error {
			return s.InitConfig()
		},
	}
}
