package square

import (
	"github.com/nickrobinson/square-cli/pkg/config"
)

type Square struct {
	*config.Config
}

func (sq *Square) InitConfig() {
	sq.Config = config.New()
}
