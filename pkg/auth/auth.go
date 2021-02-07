package auth

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/shien/weightrec-backend/pkg/authcfg"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserInfo struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

// GetLoginURL get url to login using Google account
func GetLoginURL() string {
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
	return url
}

// GetUserInfo get user infomation from google account
func GetUserInfo(code string) (string, error) {
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

	tok, err := oauth2conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Println("Filed Exchange:", err)
		return "", err
	}

	client := oauth2conf.Client(oauth2.NoContext, tok)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Failed get userinfo from Google:", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Readall failed:", err)
		return "", err
	}

	data := base64.StdEncoding.EncodeToString(body)

	return data, nil
}

func GetMailAddress(b64userinfo string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(b64userinfo)
	if err != nil {
		log.Println("Base64 Decode failed:", err)
		return "", err
	}

	var u UserInfo
	if err := json.Unmarshal(data, &u); err != nil {
		log.Println("Unmarshal failed:", err)
		return "", err
	}

	return u.Email, nil
}
