package initializer

import (
	"golang.org/x/oauth2"
	"os"
)

var (
	GithubOauthConfig *oauth2.Config
	GithubState       string
)

func InitGithubOAuth() {
	GithubOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GITHUB_CLIENT_ID"),
		ClientSecret: os.Getenv("GITHUB_SECRET"),
		Scopes:       []string{"user:email"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://github.com/login/oauth/authorize",
			TokenURL: "https://github.com/login/oauth/access_token",
		},
	}
	GithubState = os.Getenv("GITHUB_STATE_SECRET")
}
