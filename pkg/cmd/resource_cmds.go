package cmd

import (
	"net/http"

	"github.com/nickrobinson/square-cli/pkg/cmd/resource"
	"github.com/spf13/cobra"
)

func addAllResourceCmds(rootCmd *cobra.Command) {
	// Resource Commands
	customersCmd := resource.NewResourceCmd(rootCmd, "customers")

	resource.NewOperationCmd(customersCmd.Cmd, "list", "/v2/customers", http.MethodGet, map[string]string{
		"cursor":     "string",
		"sort_field": "string",
		"sort_order": "string",
	})

	resource.NewOperationCmd(customersCmd.Cmd, "delete", "/v2/customers/{customer_id}", http.MethodDelete, map[string]string{})
}
