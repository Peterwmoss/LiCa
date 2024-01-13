package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type (
	UserInfo struct {
		Id            string `json:"id"`
		Email         string `json:"email"`
		VerifiedEmail bool   `json:"verified_email"`
		Picture       string `json:"picture"`
		Hd            string `json:"hd"`
	}
)

func GetUserInfo(token string) (*UserInfo, error) {
  log.Debug().Msgf("Fetching userinfo with token: %s", token)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		return nil, err
	}

  if resp.StatusCode < 200 || resp.StatusCode > 300 {
    return nil, errors.New(fmt.Sprintf("Status code from userinfo response: %d", resp.StatusCode))
  }

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

  log.Debug().Msgf("Raw user info: %s", string(userInfoBytes))

	userInfo := &UserInfo{}
	err = json.Unmarshal(userInfoBytes, userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, err
}
