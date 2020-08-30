package main

import (
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
	"net/http"
)

func main() {
	cmd := buildRootCmd()
	cmd.Execute()
}

func buildRootCmd() *cobra.Command {
	sq := square.New()
	cmd := &cobra.Command{
		Use:   "square",
		Short: "The official command-line tool to interact with Square.",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Annotations: make(map[string]string),
	}

	// Customers Resource
	rCustomersCmd := buildResourceCommand(sq, "customers")
	rCustomersCmd.AddCommand(buildOperationCommand(sq, "list", "/v2/customers", http.MethodGet, map[string]string{
		"cursor":     "string",
		"sort_field": "string",
		"sort_order": "string",
	}))
	rCustomersCmd.AddCommand(buildOperationCommand(sq, "delete", "/v2/customers/{customer_id}", http.MethodDelete, map[string]string{}))
	rCustomersCmd.AddCommand(buildOperationCommand(sq, "get", "/v2/customers/{customer_id}", http.MethodGet, map[string]string{}))

	rInvoicesCmd := buildResourceCommand(sq, "invoices")
	rInvoicesCmd.AddCommand(buildOperationCommand(sq, "list", "/v2/invoices", http.MethodGet, map[string]string{
		"cursor":      "string",
		"sort_field":  "string",
		"sort_order":  "string",
		"location_id": "string",
	}))

	cmd.AddCommand(rCustomersCmd)
	cmd.AddCommand(rInvoicesCmd)
	cmd.AddCommand(buildResourceCommand(sq, "customer-groups"))

	cmd.AddCommand()

	return cmd
}
