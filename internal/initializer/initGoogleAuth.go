package initializer

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

var (
	GoogleOauthConfig *oauth2.Config
	GoogleState       string
)

func InitGoogleOAuth() {
	// OAuth configuration
	GoogleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("OAUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("OAUTH_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "https://forum.flodev.tech:3000/auth/google/callback",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	}
	GoogleState = os.Getenv("OAUTH_STATE_SECRET")
}
