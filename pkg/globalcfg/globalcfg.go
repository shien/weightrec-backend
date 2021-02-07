package globalcfg

import (
	"log"

	"github.com/joeshaw/envdecode"
)

// GlobalConfig is used from all scope.
type GlobalConfig struct {
	Domain string `env:"DOMAIN,required"`
}

// GetDomain return this App domain
func GetDomain() string {
	var conf GlobalConfig

	if err := envdecode.Decode(&conf); err != nil {
		log.Fatalln(err)
	}
	return conf.Domain
}
