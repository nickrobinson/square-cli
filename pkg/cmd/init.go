package cmd

import (
	"github.com/nickrobinson/square-cli/pkg/validators"
	"github.com/spf13/cobra"
)

type initCmd struct {
	Cmd *cobra.Command
}

func newInitCmd() *initCmd {
	ic := &initCmd{}

	ic.Cmd = &cobra.Command{
		Use:   "init",
		Args:  validators.ExactArgs(1),
		Short: "Initialize Square CLI config.",
		Long:  `Initialize Square CLI configuration file.`,

		RunE: ic.runInitCmd,
	}

	return ic
}

func (ic *initCmd) runInitCmd(cmd *cobra.Command, args []string) error {
	// TODO: Add init logic
	return nil
}
