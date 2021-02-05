package square

import (
	"github.com/spf13/cobra"
)

func (s *Square) GetRequest(cmd *cobra.Command, args []string) error {
	s.RequestConfig.Method = "GET"
	return s.RequestConfig.RunRequestsCmd(cmd, args)
}

func (s *Square) PutRequest(cmd *cobra.Command, args []string) error {
	s.RequestConfig.Method = "PUT"
	return s.RequestConfig.RunRequestsCmd(cmd, args)
}

func (s *Square) PostRequest(cmd *cobra.Command, args []string) error {
	s.RequestConfig.Method = "POST"
	return s.RequestConfig.RunRequestsCmd(cmd, args)
}

func (s *Square) DeleteRequest(cmd *cobra.Command, args []string) error {
	s.RequestConfig.Method = "DELETE"
	return s.RequestConfig.RunRequestsCmd(cmd, args)
}
