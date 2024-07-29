package auth

import (
	"fmt"
	"log/slog"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	TokenCookieName        = "token"
	RefreshTokenCookieName = "token_refresh"
	TokenExpiryCookieName  = "token_expiry"
)

func NewOauth2Config(baseUrl string) *oauth2.Config {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	if clientId == "" {
		slog.Warn("Client ID not specified. Please set env variable GOOGLE_CLIENT_ID")
	}

	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	if clientSecret == "" {
		slog.Warn("Client Secret not specified. Please set env variable GOOGLE_CLIENT_SECRET")
	}

	host := os.Getenv("LICA_HOST")
	if host == "" {
		host = "http://localhost:3000"
	}
	slog.Debug(fmt.Sprintf("Host: %s", host))

	return &oauth2.Config{
		RedirectURL:  host + baseUrl + "/callback",
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}
