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
	slog.Debug(fmt.Sprintf("ClientID: %s", clientId))

	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	slog.Debug(fmt.Sprintf("ClientSecret: %s", clientSecret))

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
