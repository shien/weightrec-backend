package auth

import (
	"github.com/shien/weightrec-backend/pkg/authcfg"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthInfo struct {
	AuthConfig oauth2.Config
	LoginURL   string
}

func InitAuthConf() (string, error) {
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

func CallBack() string {
	return ""
}
