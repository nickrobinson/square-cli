package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nickrobinson/square-cli/pkg/requests"
	"github.com/nickrobinson/square-cli/pkg/validators"
)

type deleteCmd struct {
	reqs requests.Base
}

func newDeleteCmd() *deleteCmd {
	gc := &deleteCmd{}

	gc.reqs.Method = "DELETE"
	gc.reqs.Cmd = &cobra.Command{
		Use:   "delete",
		Args:  validators.ExactArgs(1),
		Short: "Make DELETE requests to the Square API.",
		Long: `Make DELETE requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Delete a customer:
$ square delete /v2/customers/CGQ7M5073H2RQABDMCJDCX7RF4`,

		RunE: gc.reqs.RunRequestsCmd,
	}

	gc.reqs.InitFlags()

	return gc
}
