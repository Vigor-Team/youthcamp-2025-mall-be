package mtl

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"

	"github.com/Vigor-Team/youthcamp-2025-mall-be/app/gateway/utils"
	"github.com/cloudwego/hertz/pkg/route"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
)

var TracerProvider *tracesdk.TracerProvider

func InitTracing() route.CtxCallback {
	exporter, err := otlptracegrpc.New(context.Background())
	if err != nil {
		panic(err)
	}
	processor := tracesdk.NewBatchSpanProcessor(exporter)
	res, err := resource.New(context.Background(), resource.WithAttributes(semconv.ServiceNameKey.String(utils.ServiceName)))
	if err != nil {
		hlog.Errorf("failed to create resource: %v", err)
		res = resource.Default()
	}
	TracerProvider = tracesdk.NewTracerProvider(tracesdk.WithSpanProcessor(processor), tracesdk.WithResource(res))
	otel.SetTracerProvider(TracerProvider)

	return func(ctx context.Context) {
		exporter.Shutdown(ctx) //nolint:errcheck
	}
}
