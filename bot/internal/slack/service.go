package slack

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// IService defines the interface for the Slack bot service
type IService interface {
	HandleQuery(userID, channelID, text string)
}

type BotService struct {
	Client       *slack.Client
	SocketClient *socketmode.Client
	BackendURL   string
}

func NewService(cfg BotConfig) IService {
	return &BotService{
		Client:       cfg.Client,
		SocketClient: cfg.SocketClient,
		BackendURL:   cfg.BackendURL,
	}
}

func (b *BotService) HandleQuery(userID, channelID, text string) {
	log.Printf("Handling query from %s: %s", userID, text)

	req := ChatRequest{
		UserID: userID,
		Query:  text,
	}

	resp, err := http.Post(
		fmt.Sprintf("%s/v1/chat/stream", b.BackendURL),
		"application/json",
		ToReader(req),
	)
	if err != nil {
		log.Printf("Failed to contact backend: %v", err)
		b.sendMessage(channelID, "Backend unavailable. Please try again later.")
		return
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	for {
		var chunk StreamResponse
		if err := decoder.Decode(&chunk); err == io.EOF {
			break
		} else if err != nil {
			log.Printf("Error decoding stream chunk: %v", err)
			continue
		}

		if chunk.TextChunk != "" {
			log.Printf("Sending chunk: %s", chunk.TextChunk)
			b.sendMessage(channelID, chunk.TextChunk)
		}
	}
}

func (b *BotService) sendMessage(channelID, msg string) {
	_, _, err := b.Client.PostMessage(channelID, slack.MsgOptionText(msg, false))
	if err != nil {
		log.Printf("Failed to send message to Slack: %v", err)
	}
}
