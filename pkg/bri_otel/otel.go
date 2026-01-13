package bri_otel

import (
	"context"
	"log"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"

	"go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.37.0"
)

func NewOtelClient(timeout int) *http.Client {
	return &http.Client{
		Timeout:   time.Second * time.Duration(timeout),
		Transport: otelhttp.NewTransport(http.DefaultTransport), // need this to propagate via context.
	}
}

// otel resource
func Resource(v string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL, // pointless?
		semconv.ServiceName("briwagner/entApi"),
		semconv.ServiceVersion(v),
	)
}

func SetupOtel(ctx context.Context, fullName, version, appName string) (*sdklog.LoggerProvider, *sdktrace.TracerProvider) {
	exporter, err := otlploggrpc.New(ctx)
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	lp := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(sdklog.NewBatchProcessor(exporter)),
		sdklog.WithResource(Resource(version)),
	)

	exporterTrace, err := otlptrace.New(ctx, otlptracehttp.NewClient())
	if err != nil {
		log.Println(err)
		return nil, nil
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()), // consider lower for prod traffic
		sdktrace.WithBatcher(exporterTrace),
		sdktrace.WithResource(Resource(version)),
	)

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		))
	return lp, tp
}
