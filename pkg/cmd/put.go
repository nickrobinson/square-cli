package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nickrobinson/square-cli/pkg/requests"
	"github.com/nickrobinson/square-cli/pkg/validators"
)

type putCmd struct {
	reqs requests.Base
}

func newPutCmd() *putCmd {
	gc := &putCmd{}

	gc.reqs.Method = "PUT"
	gc.reqs.Cmd = &cobra.Command{
		Use:   "put",
		Args:  validators.ExactArgs(1),
		Short: "Make PUT requests to the Square API.",
		Long: `Make PUT requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Update a customer:
$ square put /v2/customers/CGQ7M5073H2RQABDMCJDCX7RF4 -d company_name=Square`,

		RunE: gc.reqs.RunRequestsCmd,
	}

	gc.reqs.InitFlags(true)

	return gc
}
