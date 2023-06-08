package initializer

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var (
	OauthConfig *oauth2.Config
	State       string
)

func InitOAuth() {
	// OAuth configuration
	OauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	State = os.Getenv("OAUTH_STATE_SECRET")
}
