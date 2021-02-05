package square

import (
	"github.com/nickrobinson/square-cli/internal/requests"
	"github.com/nickrobinson/square-cli/pkg/config"
)

type Square struct {
	Config        *config.Config
	RequestConfig *requests.Base
}

func New() *Square {
	c := config.Config{}
	requestConfig := requests.Base{Config: &c}
	s := &Square{
		Config:        &c,
		RequestConfig: &requestConfig,
	}
	return s
}
