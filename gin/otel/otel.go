package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/zipkin"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

var Tracer trace.Tracer

func InitTracer(serviceName string) {
	ctx := context.Background()

	exporter, err := zipkin.New("http://localhost:9411/api/v2/spans")
	if err != nil {
		zap.L().Warn("zipkin exporter not available, tracing disabled", zap.Error(err))
		Tracer = otel.Tracer(serviceName)
		return
	}

	resources, err := resource.New(ctx,
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
		),
	)
	if err != nil {
		zap.L().Error("failed to create resource", zap.Error(err))
		Tracer = otel.Tracer(serviceName)
		return
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(resources),
	)

	otel.SetTracerProvider(tp)
	Tracer = tp.Tracer(serviceName)
	zap.L().Info("opentelemetry tracer initialized")
}
