package square

import (
	"github.com/nickrobinson/square-cli/internal/requests"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func (s *Square) GetRequest(cmd *cobra.Command, args []string) error {
	log.Info("Running GetRequest")
	getRequest := requests.Base{Method: "GET", Config: s.Config}
	return getRequest.RunRequestsCmd(cmd, args)
}

func (s *Square) PutRequest(cmd *cobra.Command, args []string) error {
	putRequest := requests.Base{Method: "PUT", Config: s.Config}
	return putRequest.RunRequestsCmd(cmd, args)
}

func (s *Square) PostRequest(cmd *cobra.Command, args []string) error {
	postRequest := requests.Base{Method: "POST", Config: s.Config}
	return postRequest.RunRequestsCmd(cmd, args)
}

func (s *Square) DeleteRequest(cmd *cobra.Command, args []string) error {
	deleteRequest := requests.Base{Method: "DELETE", Config: s.Config}
	return deleteRequest.RunRequestsCmd(cmd, args)
}
