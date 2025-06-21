# 🛰️ ChatRelay - High Performance Slack Bot with OpenTelemetry

ChatRelay is a production-grade Slack bot written in **Golang**, designed to:

- 🔁 Listen for messages on Slack via Socket Mode  
- 🚀 Forward queries to a backend (real or mock)  
- 🔄 Stream responses back to the user in real-time  
- 🔍 Use **OpenTelemetry** for full observability (logs, traces)  

> Built with scalability, observability, and concurrency in mind.

---

## 📦 Repositories

| Repo | Purpose |
|------|---------|
| [chat-relay](https://github.com/Mihir99-mk/chat-relay) | Main bot source code |
| [chat-relay-lib](https://github.com/Mihir99-mk/chat-relay-lib) | Shared models, errors, and utility logic |

---

## 📐 Architecture

```
Slack ←→ ChatRelay (Golang) ←→ Chat Backend (Mock or Real)
                      ↑
          OpenTelemetry Logs + Traces
```

### 🧩 Modules

- `auth/`: Handles Slack OAuth login + callback  
- `bot/`: Manages socketmode and Slack events  
- `config/`: Loads environment configuration  
- `telemetry/`: OpenTelemetry trace and logging setup  
- `mock/`: SSE-based mock chat backend (for testing)  
- `lib/`: Custom error types, common structs  

---

## ⚙️ Configuration & Environment Variables

| Variable | Description |
|----------|-------------|
| `SLACK_BOT_TOKEN` | Slack Bot Token (xoxb-...) |
| `SLACK_APP_TOKEN` | Slack App Token (xapp-...) |
| `BACKEND_URL` | URL for the chat backend (e.g., `http://localhost:8081`) |
| `OTLP_ENDPOINT` | OpenTelemetry OTLP endpoint (e.g., `http://localhost:4317`) |
| `PORT` | Port to run the bot (default: `3000`) |

Create a `.env` file or use system environment variables.

---

## 🚀 How to Run

### 🧪 Local Run

```bash
# Run mock backend
go run main.go

# Run chat relay bot
go run main.go
```

### 🐳 Docker

```bash
docker build -t chatrelay .
docker run -p 3000:3000 --env-file .env chatrelay
```

Use `docker-compose` if running mock backend and OpenTelemetry Collector too.

---

## 🔒 Slack App Setup

1. Go to [Slack API](https://api.slack.com/apps)  
2. Create a new app & enable **Socket Mode**  
3. Add OAuth scopes:  
   - `app_mentions:read`, `chat:write`, `channels:history`, `im:history`  
4. Set redirect URL: `http://localhost:3000/api/auth/slack/callback`  
5. Copy:
   - **Bot Token** (`xoxb-...`)
   - **App Token** (`xapp-...`)  
6. Install to your workspace  

---

## 🔬 Observability (OpenTelemetry)

This bot uses **OpenTelemetry Go SDK** to generate structured logs and distributed traces.

| Feature | Description |
|---------|-------------|
| ✅ Logs | Structured `otelzap` logs with `trace_id`, `span_id` |
| ✅ Traces | Full trace from Slack message → Backend → Slack response |
| ✅ Exporter Config | Supports OTLP → Jaeger, Loki, Grafana, etc. |

Use this collector config:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
exporters:
  logging:
    loglevel: debug
  jaeger:
    endpoint: "localhost:14250"
    insecure: true
service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [jaeger, logging]
```

---

## 🧪 Testing

```bash
go test ./... -v
```

- ✅ Unit tests: Slack handlers, backend client, error handling  
- ✅ Integration test: Slack events → Mock backend  
- 📁 Tests live inside `test/` or co-located with components  

---

## 🛍️ Slack App Marketplace Plan

The app is ready for Slack Marketplace submission with:  
- OAuth 2.0 flow  
- Proper scopes and manifest  
- Security best practices (token encryption, input validation)  
- Logging, tracing, and diagnostics  

---

## ✨ Highlights

- Built using idiomatic Go  
- Clean and testable architecture  
- SocketMode Slack bot with concurrent request handling  
- Custom middleware and structured errors  
- Fully observable via OpenTelemetry  

---

## 📄 License

MIT License © Mihir Khode
