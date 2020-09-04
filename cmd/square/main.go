package main

import (
	"github.com/nickrobinson/square-cli/pkg/config"
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net/http"
)

func main() {
	cmd := buildRootCmd()
	cmd.Execute()
}

var (
	sq square.Square
)

func buildRootCmd() *cobra.Command {
	cobra.OnInitialize(initConfig)

	rootCmd := &cobra.Command{
		Use:   "square",
		Short: "The official command-line tool to interact with Square.",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Annotations: make(map[string]string),
	}

	rootCmd.PersistentFlags().StringP("profile", "p", "default", "profile to use")
	viper.BindPFlag("profile", rootCmd.PersistentFlags().Lookup("profile"))

	// Customers Resource
	rCustomersCmd := buildResourceCommand(&sq, "customers")
	rCustomersCmd.AddCommand(buildOperationCommand(&sq, "list", "/v2/customers", http.MethodGet, map[string]string{
		"cursor":     "string",
		"sort_field": "string",
		"sort_order": "string",
	}))
	rCustomersCmd.AddCommand(buildOperationCommand(&sq, "delete", "/v2/customers/{customer_id}", http.MethodDelete, map[string]string{}))
	rCustomersCmd.AddCommand(buildOperationCommand(&sq, "get", "/v2/customers/{customer_id}", http.MethodGet, map[string]string{}))

	rInvoicesCmd := buildResourceCommand(&sq, "invoices")
	rInvoicesCmd.AddCommand(buildOperationCommand(&sq, "list", "/v2/invoices", http.MethodGet, map[string]string{
		"cursor":      "string",
		"sort_field":  "string",
		"sort_order":  "string",
		"location_id": "string",
	}))

	rootCmd.AddCommand(rCustomersCmd)
	rootCmd.AddCommand(rInvoicesCmd)
	rootCmd.AddCommand(buildResourceCommand(&sq, "customer-groups"))

	return rootCmd
}

func initConfig() {
	sq.Config = config.New()
}
