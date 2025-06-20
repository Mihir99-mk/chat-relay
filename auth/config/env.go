package config

import (
	"log"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
)

type IEnv interface {
	Get(key string) string
	GetPort() string
	GetDBPort() string
	GetDBUsername() string
	GetDBPassword() string
	GetDBHost() string
	GetDBName() string
	GetSlackClientId() string
	GetSlackSecret() string
	GetSlackRedirectUrl() string
	GetServiceName() string
	GetOTLPEndpoint() string
	GetDaprHttpPort() string
	GetDaprGrpcPort() string
}

var envOnce sync.Once

type env struct{}

func NewEnv() IEnv {
	envOnce.Do(func() {
		envFile := ".env"

		err := godotenv.Load(envFile)
		if err != nil {
			log.Fatalf("Error loading .env file from %s: %s", envFile, err)
		}

		log.Println(".env file loaded successfully")
	})
	return &env{}
}

const (
	DB_PORT             = "DB_PORT"
	PORT                = "PORT"
	DB_USERNAME         = "DB_USERNAME"
	DB_PASSWORD         = "DB_PASSWORD"
	DB_HOST             = "DB_HOST"
	DB_NAME             = "DB_NAME"
	SLACK_CLIENT_ID     = "SLACK_CLIENT_ID"
	SLACK_CLIENT_SECRET = "SLACK_CLIENT_SECRET"
	SLACK_REDIRECT_URL  = "SLACK_REDIRECT_URL"

	SERVICE_NAME   = "SERVICE_NAME"
	OTLP_ENDPOINT  = "OTLP_ENDPOINT"
	DAPR_HTTP_PORT = "DAPR_HTTP_PORT"
	DAPR_GRPC_PORT = "DAPR_GRPC_PORT"
)

func (e *env) Get(key string) string {
	return strings.TrimSpace(os.Getenv(key))
}

func (e *env) GetPort() string {
	return os.Getenv(PORT)
}

func (e *env) GetDBPort() string {
	return os.Getenv(DB_PORT)
}

func (e *env) GetDBPassword() string {
	return os.Getenv(DB_PASSWORD)
}

func (e *env) GetDBUsername() string {
	return os.Getenv(DB_USERNAME)
}

func (e *env) GetDBHost() string {
	return os.Getenv(DB_HOST)
}

func (e *env) GetDBName() string {
	return os.Getenv(DB_NAME)
}

func (e *env) GetSlackClientId() string {
	return os.Getenv(SLACK_CLIENT_ID)
}

func (e *env) GetSlackRedirectUrl() string {
	return os.Getenv(SLACK_REDIRECT_URL)
}

func (e *env) GetSlackSecret() string {
	return os.Getenv(SLACK_CLIENT_SECRET)
}

func (e *env) GetServiceName() string {
	return os.Getenv(SERVICE_NAME)
}

func (e *env) GetOTLPEndpoint() string {
	return os.Getenv(OTLP_ENDPOINT)
}

func (e *env) GetDaprGrpcPort() string {
	return os.Getenv(DAPR_GRPC_PORT)
}

func (e *env) GetDaprHttpPort() string {
	return os.Getenv(DAPR_HTTP_PORT)
}
