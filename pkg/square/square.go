package square

import "github.com/nickrobinson/square-cli/pkg/config"

type Square struct {
	Config *config.Config
}

func New() *Square {
	c, err := config.Profile
	if err != nil {
		return nil
	}
	s := &Square{
		Config: c,
	}
	return s
}
