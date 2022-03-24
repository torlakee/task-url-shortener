package deployment

import (
	"github.com/couchbase/gocb/v2"
)

type Application struct {
	Storage  *gocb.Bucket
	Title string
	Configuration *AppConf
}

type AppConf struct {
	External struct {
		Host string
	}
	Internal struct {
		Server struct {
			Port int
		}
		Database struct {
			Host string
			Storage string
			User string
			Password string
			Port uint16
		}
	}
}
