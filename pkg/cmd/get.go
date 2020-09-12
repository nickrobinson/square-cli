package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nickrobinson/square-cli/pkg/requests"
	"github.com/nickrobinson/square-cli/pkg/validators"
)

type getCmd struct {
	reqs requests.Base
}

func newGetCmd() *getCmd {
	gc := &getCmd{}

	gc.reqs.Method = "GET"
	gc.reqs.Cmd = &cobra.Command{
		Use:   "get",
		Args:  validators.ExactArgs(1),
		Short: "Make GET requests to the Square API.",
		Long: `Make GET requests to the Square API.

You can only get data in test mode, the get command does not work for live mode.
The command also supports common API features like pagination and limits.

For a full list of supported paths, see the API reference: https://square.com/docs/api

GET a charge:
$ square get /invoices/1V2QJAFF28W8N58G0K3FBNDYTW

GET 50 charges:
$ square get --limit 50 /invoices`,

		RunE: gc.reqs.RunRequestsCmd,
	}

	gc.reqs.InitFlags(true)

	return gc
}
