package main

import (
	"github.com/nickrobinson/square-cli/pkg/square"
	"github.com/nickrobinson/square-cli/pkg/validators"
	"github.com/spf13/cobra"
)

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
	initFlags(cmd)
	return cmd
}

func buildPutCommand(s *square.Square) *cobra.Command {
	cmd := cobra.Command{
		Use:   "put",
		Args:  validators.ExactArgs(1),
		Short: "Make PUT requests to the Square API.",
		Long: `Make PUT requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Update a customer:
$ square put /v2/customers/CGQ7M5073H2RQABDMCJDCX7RF4 -d '{"company_name": "Square"}'`,
		RunE: s.PutRequest,
	}
	return &cmd
}

func buildPostCommand(s *square.Square) *cobra.Command {
	cmd := cobra.Command{
		Use:   "post",
		Args:  validators.ExactArgs(1),
		Short: "Make POST requests to the Square API.",
		Long: `Make POST requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Update a customer:
$ square post /v2/customers -d '{"given_name": "Jack", "family_name": "Dorsey"}'`,
		RunE: s.PostRequest,
	}
	return &cmd
}

func buildDeleteCommand(s *square.Square) *cobra.Command {
	cmd := cobra.Command{
		Use:   "delete",
		Args:  validators.ExactArgs(1),
		Short: "Make DELETE requests to the Square API.",
		Long: `Make DELETE requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Delete a customer:
$ square delete /v2/customers/CGQ7M5073H2RQABDMCJDCX7RF4`,
		RunE: s.DeleteRequest,
	}
	return &cmd
}
