package cmd

import (
	"github.com/spf13/cobra"

	"github.com/nickrobinson/square-cli/pkg/requests"
	"github.com/nickrobinson/square-cli/pkg/validators"
)

type postCmd struct {
	reqs requests.Base
}

func newPostCmd() *postCmd {
	gc := &postCmd{}

	gc.reqs.Method = "POST"
	gc.reqs.Cmd = &cobra.Command{
		Use:   "post",
		Args:  validators.ExactArgs(1),
		Short: "Make POST requests to the Square API.",
		Long: `Make POST requests to the Square API.

For a full list of supported paths, see the API reference: https://developer.squareup.com/reference/square

Update a customer:
$ square post /v2/customers -d '{"given_name": "Jack", "family_name": "Dorsey"}'`,

		RunE: gc.reqs.RunRequestsCmd,
	}

	gc.reqs.InitFlags()

	return gc
}