package tracing

import (
	"context"

	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"google.golang.org/grpc"
)

func NewOTLPExporter(ctx context.Context, endpoint string) *otlptrace.Exporter {
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),          // non TLS
		otlptracegrpc.WithEndpoint(endpoint),  // e.g: "tempo:4317"
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)

	exporter, err := otlptrace.New(ctx, client)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create OTLP trace exporter")
	}
	return exporter
}
