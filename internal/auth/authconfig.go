package auth

import (
	"github.com/rs/zerolog/log"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

type (
	AuthConfig struct {
    BaseUrl string
  }
)

func NewAuthConfig(baseUrl string) *AuthConfig {
  return &AuthConfig{ baseUrl }
}

func (ac *AuthConfig) Get() *oauth2.Config {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	log.Debug().Msgf("ClientID: %s", clientId)

	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	log.Debug().Msgf("ClientSecret: %s", clientSecret)

	host := os.Getenv("LICA_HOST")
	if host == "" {
		host = "http://localhost:3000"
	}
	log.Debug().Msgf("Host: %s", host)

	return &oauth2.Config{
		RedirectURL:  host + ac.BaseUrl + "/callback",
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
