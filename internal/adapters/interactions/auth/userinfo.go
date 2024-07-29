package auth

import (
	"context"
	"encoding/json"
	"fmt"
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
		return UserInfo{}, fmt.Errorf("auth.GetUserInfo: Get request failed:\n%w", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode > 300 {
		return UserInfo{}, fmt.Errorf("auth.GetUserInfo: Response status code is non-success:\n%v", resp)
	}

	userInfoBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return UserInfo{}, fmt.Errorf("auth.GetUserInfo: Failed to read response body:\n%w", err)
	}

	userInfo := UserInfo{}
	err = json.Unmarshal(userInfoBytes, &userInfo)
	if err != nil {
		return UserInfo{}, fmt.Errorf("auth.GetUserInfo: Failed to unmarshal:\n%w", err)
	}

	return userInfo, nil
}
