package auth

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type (
	UserInfo struct {
		Id            string
		Email         string
		VerifiedEmail bool
		Picture       string
		Hd            string
	}
)

func GetUserInfo(token string) (*UserInfo, error) {
	log.Debug().Msgf("Fetching userinfo with token: %s", token)
	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token)
	if err != nil {
		log.Error().Err(err).Msg("Failed to get:")
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		err = errors.New("Non success response")
		log.Error().Err(err).Msgf("Status code from userinfo response: %d", resp.StatusCode)
		return nil, err
	}

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Failed to read response body:")
		return nil, err
	}

	userInfo := &UserInfo{}
	err = json.Unmarshal(userInfoBytes, userInfo)
	if err != nil {
		log.Error().Err(err).Msg("Failed to unmarshal:")
		return nil, err
	}

	return userInfo, nil
}
