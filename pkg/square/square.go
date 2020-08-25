package square

import (
	"github.com/nickrobinson/square-cli/pkg/config"
)

type Square struct {
	*config.Config
}

func New() *Square {
	c := config.New()
	return NewWithConfig(c)
}

func NewWithConfig(c *config.Config) *Square {
	return &Square{
		Config: c,
	}
}
