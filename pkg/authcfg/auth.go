package authcfg

import (
	"log"

	"github.com/joeshaw/envdecode"
)

// CredConfig has Credentials for Google auth
type CredConfig struct {
	ClientID     string `env:"CLIENT_ID,required"`
	ClientSecret string `env:"CLIENT_SECRET,required"`
	RedirectURL  string `env:"REDIRECT_URL,required"`
}

// GetCredConf return Credentials configuration
func GetCredConf() *CredConfig {
	var conf CredConfig

	if err := envdecode.Decode(&conf); err != nil {
		log.Fatalln(err)
	}

	return &conf
}
