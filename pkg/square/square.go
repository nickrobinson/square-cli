package square

import (
	"github.com/nickrobinson/square-cli/pkg/config"
	log "github.com/sirupsen/logrus"
)

type Square struct {
	Config *config.Config
}

func New() *Square {
	c := config.Config{}
	err := c.Load()
	if err != nil {
		log.Error(err)
		return nil
	}
	s := &Square{
		Config: &c,
	}
	return s
}
