package square

import (
	"fmt"

	"github.com/nickrobinson/square-cli/internal/status"
)

func (s *Square) GetSquareStatus() error {
	status, err := status.GetSquareStatus()
	if err != nil {
		return err
	}
	fmt.Println(status)
	return nil
}
