package event

import (
	"bot/config"
	"log"

	cfg "bot/internal/slack"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

func InitSlackEvent() {
	env := config.NewEnv()
	botToken := env.GetSlackBotToken()
	appToken := env.GetSlackAppToken()
	backendURL := ""

	api := slack.New(
		botToken,
		slack.OptionDebug(true),
		slack.OptionAppLevelToken(appToken),
	)

	socketClient := socketmode.New(
		api,
		socketmode.OptionDebug(true),
	)

	bot := &cfg.BotConfig{
		Client:       api,
		SocketClient: socketClient,
		BackendURL:   backendURL,
	}

	handler := cfg.NewHandler(*bot)
	go handler.HandleEvents(socketClient)

	log.Println("ChatRelay bot running...")
	err := socketClient.Run()
	if err != nil {
		log.Println("error occur in socket client")
	}
}
