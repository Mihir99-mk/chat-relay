package auth

import (
	"auth/config"
	"auth/util/common"
	"context"
	"log"

	"github.com/Mihir99-mk/chat-relay-lib/wrapperx"
	"github.com/slack-go/slack"
)

type IService interface {
	SlackCallback(ctx context.Context, code string) error
}

type service struct {
	storage IStorage
	env     config.IEnv
}

func NewService(config config.IConfig) IService {
	return &service{
		storage: NewStorage(config),
		env:     config.Env(),
	}
}

func (s *service) SlackCallback(ctx context.Context, code string) error {
	token, err := common.ExchangeCodeForToken(code, s.env)
	if err != nil {
		log.Printf("Failed to exchange token: %v", err)
		return wrapperx.WrapToOauthError(Domain, err)
	}

	api := slack.New(token.AccessToken)
	authTest, err := api.AuthTest()
	if err != nil {
		log.Printf("Auth test failed: %v", err)
		return wrapperx.WrapToOauthError(Domain, err)
	}

	userInfo, err := api.GetUserInfo(authTest.UserID)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return wrapperx.WrapToCustomError(Domain, err)
	}

	provider, err := s.storage.GetAuthProvider(ctx)
	if err != nil {
		return wrapperx.WrapToEntError(Domain, err)
	}

	slackUser := &SlackUser{
		ProviderID:     provider.ID,
		ProviderUserID: userInfo.ID,
		TeamID:         authTest.TeamID,
		Name:           userInfo.Name,
		RealName:       userInfo.RealName,
		Email:          userInfo.Profile.Email,
		AccessToken:    token.AccessToken,
		Scope:          "",
		RawProfile:     nil,
	}

	_, err = s.storage.UpsertSlackUser(ctx, slackUser)
	if err != nil {
		log.Printf("Failed to save Slack user: %v", err)
		return wrapperx.WrapToEntError(Domain, err)
	}

	return nil
}
