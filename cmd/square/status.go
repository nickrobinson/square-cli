package main

import (
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
)

func buildStatusCommand(s *square.Square) *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Args:  cobra.NoArgs,
		Short: "Check the operational status of Square.",
		Long:  `Check the operational status of Square.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return s.GetSquareStatus()
		},
	}
}
