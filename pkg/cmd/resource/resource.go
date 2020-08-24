package resource

import (
	"github.com/spf13/cobra"
)

//
// Public types
//

// ResourceCmd represents Resource commands. Resource commands can be either
// top-level commands or nested under namespace commands. Resource commands
// are containers for operation commands.
//
// Example of Resources: `customers`, `payment_intents` (top-level, not
// namespaced), `early_fraud_warnings` (namespaced under `radar`).
type ResourceCmd struct { //nolint:golint
	Cmd  *cobra.Command
	Name string
}

//
// Public functions
//

// NewResourceCmd returns a new ResourceCmd.
func NewResourceCmd(parentCmd *cobra.Command, resourceName string) *ResourceCmd {
	cmd := &cobra.Command{
		Use:         resourceName,
		Annotations: make(map[string]string),
	}

	parentCmd.AddCommand(cmd)
	parentCmd.Annotations[resourceName] = "Resource"

	return &ResourceCmd{
		Cmd:  cmd,
		Name: resourceName,
	}
}
