package slack

import (
	"log"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
)

type Handler interface {
	HandleEvents(client *socketmode.Client)
}

type SlackHandler struct {
	Bot IService
}

func NewHandler(bot BotConfig) Handler {
	return &SlackHandler{
		Bot: NewService(bot),
	}
}

func (h *SlackHandler) HandleEvents(client *socketmode.Client) {
	for evt := range client.Events {
		switch evt.Type {

		case socketmode.EventTypeInteractive:
			log.Println("Received interactive event (not implemented)")

		case socketmode.EventTypeEventsAPI:
			event, ok := evt.Data.(slackevents.EventsAPIEvent)
			if !ok {
				log.Println("Invalid event type cast")
				continue
			}

			client.Ack(*evt.Request)

			switch ev := event.InnerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				log.Printf("Mention from user %s in channel %s: %s", ev.User, ev.Channel, ev.Text)
				go h.Bot.HandleQuery(ev.User, ev.Channel, ev.Text)
			default:
				log.Printf("Ignored inner event type: %T", ev)
			}

		default:
			log.Printf("Unhandled event type: %s", evt.Type)
		}
	}
}
