package cmd

import (
	"fmt"
	"log"

	"github.com/nickrobinson/square-cli/internal/status"
	"github.com/spf13/cobra"
)

type statusCmd struct {
	Cmd *cobra.Command
}

func newStatusCmd() *statusCmd {
	ic := &statusCmd{}

	ic.Cmd = &cobra.Command{
		Use:   "status",
		Args:  cobra.NoArgs,
		Short: "Check the operational status of Square.",
		Long:  `Check the operational status of Square.`,

		RunE: ic.runStatusCmd,
	}

	return ic
}

func (ic *statusCmd) runStatusCmd(cmd *cobra.Command, args []string) error {
	status, err := status.GetSquareStatus()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(status)
	return nil
}
