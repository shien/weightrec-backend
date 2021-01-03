package auth

import (
	"github.com/shien/weightrec-backend/pkg/authcfg"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GetLoginURL() (string, error) {
	conf := authcfg.GetCredConf()
	oauth2conf := &oauth2.Config{
		ClientID:     conf.ClientID,
		ClientSecret: conf.ClientSecret,
		RedirectURL:  conf.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
	url := oauth2conf.AuthCodeURL("state")
	return url, nil
}
