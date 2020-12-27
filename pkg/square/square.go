package square

import (
	"github.com/nickrobinson/square-cli/internal/requests"
	"github.com/nickrobinson/square-cli/pkg/config"
	log "github.com/sirupsen/logrus"
)

type Square struct {
	Config        *config.Config
	RequestConfig *requests.Base
}

func New() *Square {
	c := config.Config{}
	requestConfig := requests.Base{Config: &c}
	err := c.Load()
	if err != nil {
		log.Error(err)
		return nil
	}
	s := &Square{
		Config:        &c,
		RequestConfig: &requestConfig,
	}
	return s
}
