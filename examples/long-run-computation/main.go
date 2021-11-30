package main

import (
	"context"
	"io"
	"log"
	"os"
	"os/signal"

	"github.com/Gogistics/prj-logging-tracing-v1/examples/long-run-computation/app"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

func main() {
	logger := log.New(os.Stdout, "atai-long-computation:", 0)

	// Write telemetry data to a file
	f, err := os.Create("computation-traces.txt")
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()

	exp, err := newExporter(f)
	if err != nil {
		logger.Fatal(err)
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(newResource()),
	)

	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Fatal(err)
		}
	}()
	otel.SetTracerProvider(tp)

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	errCh := make(chan error)
	app := app.NewApp(os.Stdin, logger)
	go func() {
		errCh <- app.Run(context.Background())
	}()

	select {
	case sig := <-sigCh:
		logger.Printf("\nSignal: %v\n", sig)
		logger.Println("\nExit")
	case err := <-errCh:
		if err != nil {
			logger.Fatal(err)
		}
	}
}

func newExporter(w io.Writer) (trace.SpanExporter, error) {
	return stdouttrace.New(
		stdouttrace.WithWriter(w),
		stdouttrace.WithPrettyPrint(),
		stdouttrace.WithoutTimestamps(),
	)
}

func newResource() *resource.Resource {
	r, _ := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("atai-computation"),
			semconv.ServiceVersionKey.String("v0.0.0"),
			attribute.String("environment", "demo"),
		),
	)
	return r
}