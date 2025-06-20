package config

import (
	"context"
	"log"
	stdlog "log"

	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"go.uber.org/zap"
)

func InitTracer(ctx context.Context, env IEnv) func(context.Context) error {
	exporter, err := otlptracegrpc.New(ctx,
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(env.GetOTLPEndpoint()),
	)
	if err != nil {
		stdlog.Fatalf("failed to create trace exporter: %v", err)
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(env.GetServiceName()),
		)),
	)
	otel.SetTracerProvider(tp)
	return tp.Shutdown
}

func InitMetrics(ctx context.Context, env IEnv) func(context.Context) error {
	exporter, err := otlpmetricgrpc.New(ctx,
		otlpmetricgrpc.WithInsecure(),
		otlpmetricgrpc.WithEndpoint(env.GetOTLPEndpoint()),
	)
	if err != nil {
		stdlog.Fatalf("failed to create metrics exporter: %v", err)
	}

	mp := metric.NewMeterProvider(
		metric.WithReader(metric.NewPeriodicReader(exporter)),
		metric.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(env.GetServiceName()),
		)),
	)
	otel.SetMeterProvider(mp)
	return mp.Shutdown
}

var (
	zapLogger *otelzap.Logger
)

func InitLogger(ctx context.Context, env IEnv) func(context.Context) error {
	exporter, err := otlploggrpc.New(ctx,
		otlploggrpc.WithInsecure(),
		otlploggrpc.WithEndpoint(env.GetOTLPEndpoint()),
	)
	if err != nil {
		stdlog.Fatalf("failed to create OTLP log exporter: %v", err)
	}

	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(env.GetServiceName()),
		)),
	)

	log.Println("Using service.name:", env.GetServiceName())

	coreLogger, err := zap.NewDevelopment() // Or zap.NewProduction()
	if err != nil {
		stdlog.Fatalf("failed to create zap logger: %v", err)
	}

	zapLogger = otelzap.New(coreLogger,
		otelzap.WithLoggerProvider(loggerProvider),
	)

	otelzap.ReplaceGlobals(zapLogger)

	global.SetLoggerProvider(loggerProvider)

	stdlog.Println("Zap + OTel Logger initialized")

	return loggerProvider.Shutdown
}

func Logger() *otelzap.Logger {
	return zapLogger
}

func InitOtel(ctx context.Context) (shutdowns []func(context.Context) error) {
	env := NewEnv()
	shutdowns = append(shutdowns, InitTracer(ctx, env))
	shutdowns = append(shutdowns, InitMetrics(ctx, env))
	shutdowns = append(shutdowns, InitLogger(ctx, env))
	return
}
