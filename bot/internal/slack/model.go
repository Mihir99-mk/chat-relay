package slack

import (
	"bytes"
	"encoding/json"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

type ChatRequest struct {
	UserID string `json:"userId"`
	Query  string `json:"query"`
}

type StreamResponse struct {
	TextChunk string `json:"textChunk,omitempty"`
	Status    string `json:"status,omitempty"`
}

type FullResponse struct {
	FullResponse string `json:"fullResponse"`
}

type BotConfig struct {
	Client       *slack.Client
	SocketClient *socketmode.Client
	BackendURL   string
}

func ToReader(req ChatRequest) *bytes.Reader {
	b, _ := json.Marshal(req)
	return bytes.NewReader(b)
}
