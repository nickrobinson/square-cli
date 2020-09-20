package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nickrobinson/square-cli/internal/requests"
	"github.com/nickrobinson/square-cli/pkg/validators"
)

type getCmd struct {
	reqs requests.Base
}

func newGetCmd() *getCmd {
	gc := &getCmd{}

	gc.reqs.Method = "GET"
	gc.reqs.Profile = &Profile
	gc.reqs.Cmd = &cobra.Command{
		Use:   "get",
		Args:  validators.ExactArgs(1),
		Short: "Make GET requests to the Square API.",
		Long: `Make GET requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

GET list of customers:
$ square get customers

GET 50 invoices:
$ square get --limit 50 invoices`,

		RunE: gc.reqs.RunRequestsCmd,
	}

	gc.reqs.InitFlags()

	return gc
}
