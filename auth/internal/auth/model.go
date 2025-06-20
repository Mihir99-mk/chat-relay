package auth

import "time"

const (
	Domain = "Auth"
)

type FormData struct {
	Code  string `form:"code" validate:"required"`
	State string `form:"state" validate:"required"`
}

type SlackUser struct {
	ProviderID     int            `json:"providerId"`
	ProviderUserID string         `json:"providerUserId"`
	TeamID         string         `json:"teamId"`
	Name           string         `json:"name"`
	RealName       string         `json:"realName"`
	Email          string         `json:"email"`
	AccessToken    string         `json:"accessToken"`
	RefreshToken   string         `json:"refreshToken"`
	TokenType      string         `json:"tokenType"`
	ExpiresAt      *time.Time     `json:"expiresAt"`
	Scope          string         `json:"scope"`
	RawProfile     map[string]any `json:"rawProfile"`
}
