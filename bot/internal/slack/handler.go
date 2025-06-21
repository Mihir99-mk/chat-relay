package slack

import (
	"context"

	"github.com/slack-go/slack/slackevents"
	"github.com/slack-go/slack/socketmode"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type Handler interface {
	HandleEvents(client *socketmode.Client)
}

type SlackHandler struct {
	Bot    IService
	tracer trace.Tracer
}

func NewHandler(bot BotConfig) Handler {
	return &SlackHandler{
		Bot:    NewService(bot),
		tracer: otel.Tracer("slack/handler"),
	}
}

func (h *SlackHandler) HandleEvents(client *socketmode.Client) {
	for evt := range client.Events {
		ctx := context.Background()
		ctx, span := h.tracer.Start(ctx, "HandleSlackEvent",
			trace.WithAttributes(attribute.String("event.type", string(evt.Type))),
		)

		switch evt.Type {

		case socketmode.EventTypeInteractive:
			span.AddEvent("Received interactive event (not implemented)")

		case socketmode.EventTypeEventsAPI:
			event, ok := evt.Data.(slackevents.EventsAPIEvent)
			if !ok {
				span.RecordError(context.DeadlineExceeded)
				span.SetStatus(1, "Invalid event type cast")
				span.End()
				continue
			}

			client.Ack(*evt.Request)
			span.SetAttributes(attribute.String("slack.api_event.type", string(event.Type)))

			switch ev := event.InnerEvent.Data.(type) {
			case *slackevents.AppMentionEvent:
				span.SetAttributes(
					attribute.String("slack.user_id", ev.User),
					attribute.String("slack.channel_id", ev.Channel),
					attribute.String("slack.text", ev.Text),
				)

				go func() {
					_, childSpan := h.tracer.Start(ctx, "HandleAppMention")
					defer childSpan.End()
					h.Bot.HandleQuery(ev.User, ev.Channel, ev.Text)
				}()

			default:
				span.SetAttributes(attribute.String("slack.inner_event.type", "ignored"))
				span.AddEvent("Ignored inner event")
			}

		default:
			span.AddEvent("Unhandled event type")
		}

		span.End()
	}
}
