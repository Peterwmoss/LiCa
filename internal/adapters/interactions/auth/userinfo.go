package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log/slog"

	"golang.org/x/oauth2"
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

func GetUserInfo(ctx context.Context, token *oauth2.Token, config *oauth2.Config) (UserInfo, error) {
	slog.Debug("Fetching userinfo", "token", token)
	client := config.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		slog.Error("Get request failed", "error", err)
		return UserInfo{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		err = errors.New("Non success response")
		slog.Error("Response status code is non-success", "error", err, "response", resp)
		return UserInfo{}, err
	}

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("Failed to read response body", "error", err)
		return UserInfo{}, err
	}

	userInfo := UserInfo{}
	err = json.Unmarshal(userInfoBytes, &userInfo)
	if err != nil {
		slog.Error("Failed to unmarshal", "error", err)
		return UserInfo{}, err
	}

	return userInfo, nil
}
