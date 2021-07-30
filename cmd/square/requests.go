package main

import (
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/nickrobinson/square-cli/pkg/validators"
	"github.com/spf13/cobra"
)

func initRequestFlags(s *square.Square, cmd *cobra.Command) {
	dataUsage := "Data to pass for the API request"
	if cmd.Use == "put" || cmd.Use == "post" {
		dataUsage = "JSON data to pass in API request body"
	}
	if cmd.Use == "post" {
		cmd.Flags().StringVarP(&s.RequestConfig.Parameters.Idempotency, "idempotency", "i", "", "Sets the idempotency key for your request, preventing replaying the same requests within a 24 hour period")
	}
	cmd.Flags().StringArrayVarP(&s.RequestConfig.Parameters.Data, "data", "d", []string{}, dataUsage)
	cmd.Flags().BoolVarP(&s.RequestConfig.ShowHeaders, "show-headers", "s", false, "Show headers on responses to GET, PUT, POST, and DELETE requests")
	cmd.Flags().BoolVarP(&s.RequestConfig.AutoConfirm, "confirm", "c", false, "Automatically confirm the command being entered. WARNING: This will result in NOT being prompted for confirmation for certain commands")
	// Conditionally add flags for GET requests. I'm doing it here to keep `limit`, `start_after` and `ending_before` unexported
	if cmd.Use == "get" {
		cmd.Flags().StringVarP(&s.RequestConfig.Parameters.Limit, "limit", "l", "", "A limit on the number of objects to be returned, between 1 and 100 (default is 10)")
	}

	cmd.Flags().StringVarP(&s.RequestConfig.Parameters.Version, "api-version", "v", "", "Square API Version to use for request")

	// Hidden configuration flags, useful for dev/debugging
	cmd.Flags().StringVar(&s.RequestConfig.APIBaseURL, "base-url", "", "Sets the API base URL")
	cmd.Flags().MarkHidden("base-url")
}

func buildGetCommand(s *square.Square) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Args:  validators.ExactArgs(1),
		Short: "Make GET requests to the Square API.",
		Long: `Make GET requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

GET list of customers:
$ square get customers

GET 50 invoices:
$ square get --limit 50 invoices`,
		RunE: s.GetRequest,
	}
	initRequestFlags(s, cmd)
	return cmd
}

func buildPutCommand(s *square.Square) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "put",
		Args:  validators.ExactArgs(1),
		Short: "Make PUT requests to the Square API.",
		Long: `Make PUT requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Update a customer:
$ square put /v2/customers/CGQ7M5073H2RQABDMCJDCX7RF4 -d '{"company_name": "Square"}'`,
		RunE: s.PutRequest,
	}
	initRequestFlags(s, cmd)
	return cmd
}

func buildPostCommand(s *square.Square) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "post",
		Args:  validators.ExactArgs(1),
		Short: "Make POST requests to the Square API.",
		Long: `Make POST requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Update a customer:
$ square post /v2/customers -d '{"given_name": "Jack", "family_name": "Dorsey"}'`,
		RunE: s.PostRequest,
	}
	initRequestFlags(s, cmd)
	return cmd
}

func buildDeleteCommand(s *square.Square) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete",
		Args:  validators.ExactArgs(1),
		Short: "Make DELETE requests to the Square API.",
		Long: `Make DELETE requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Delete a customer:
$ square delete /v2/customers/CGQ7M5073H2RQABDMCJDCX7RF4`,
		RunE: s.DeleteRequest,
	}
	initRequestFlags(s, cmd)
	return cmd
}
