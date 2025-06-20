package common

import (
	"bot/config"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type SlackOAuthResponse struct {
	OK          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
	Error       string `json:"error"`
}

func ExchangeCodeForToken(code string, env config.IEnv) (*SlackOAuthResponse, error) {
	clientID := env.GetSlackClientId()
	clientSecret := env.GetSlackSecret()
	redirectURI := env.GetSlackRedirectUrl()

	resp, err := http.PostForm("https://slack.com/api/oauth.v2.access", url.Values{
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"code":          {code},
		"redirect_uri":  {redirectURI},
	})
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var oauthResp SlackOAuthResponse
	if err := json.NewDecoder(resp.Body).Decode(&oauthResp); err != nil {
		return nil, err
	}
	if !oauthResp.OK {
		return nil, fmt.Errorf("Slack OAuth error: %s", oauthResp.Error)
	}

	return &oauthResp, nil
}
