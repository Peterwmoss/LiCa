package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/Peterwmoss/LiCa/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/bun"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const stateCheck string = "a8e7hfwnkf3"

func getAuthConfig() *oauth2.Config {
	clientId := os.Getenv("GOOGLE_CLIENT_ID")
	log.Debug().Msgf("ClientID: %s", clientId)
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	log.Debug().Msgf("ClientSecret: %s", clientSecret)

	return &oauth2.Config{
		RedirectURL:  "http://localhost:3000/auth/callback",
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Scopes:       []string{"openid", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
}

type UserInfo struct {
	Id            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
	Hd            string `json:"hd"`
}

func loginHandler(authConfig *oauth2.Config) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		redirectUrl := authConfig.AuthCodeURL(stateCheck)
		return c.Status(fiber.StatusSeeOther).Redirect(redirectUrl)
	}
}

func callbackHandler(authConfig *oauth2.Config, db *bun.DB, ctx context.Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		state := c.Query("state")
		if state != stateCheck {
			log.Error().Msg("CSFR detected")
			return c.SendStatus(fiber.StatusForbidden)
		}

		code := c.Query("code")
		token, err := authConfig.Exchange(ctx, code)
		if err != nil {
			log.Error().Err(err).Msg("Failed code exchange")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		info, err := getUserInfo(token.AccessToken)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get user info from Google")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		_, err = repository.CreateUser(info.Email, db, ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create user")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: token.AccessToken,
		})
		return c.Status(fiber.StatusTemporaryRedirect).Redirect("/auth/me")
	}
}

func logoutHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		c.Cookie(&fiber.Cookie{
			Name:  "token",
			Value: "",
		})
		return c.SendStatus(fiber.StatusOK)
	}
}

func meHandler(db *bun.DB, ctx context.Context) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Cookies("token")
		if token == "" {
			return c.Status(fiber.StatusTemporaryRedirect).Redirect("/auth/login")
		}

		info, err := getUserInfo(token)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		user, err := repository.GetUser(info.Email, db, ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to create user")
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(user)
	}
}

func getUserInfo(token string) (*UserInfo, error) {
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return nil, err
	}

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	userInfo := &UserInfo{}
	err = json.Unmarshal(userInfoBytes, userInfo)
	if err != nil {
		return nil, err
	}
	return userInfo, err
}
