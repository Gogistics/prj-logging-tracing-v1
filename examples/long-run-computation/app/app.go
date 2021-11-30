package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"strconv"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	name = "Computation"
)

type App struct {
	reader io.Reader
	logger *log.Logger
}

// Create new App
func NewApp(r io.Reader, l *log.Logger) *App {
	return &App{r, l}
}

// Run: start the computation
func (a *App) Run(ctx context.Context) error {
	for {
		var span trace.Span
		ctx, span = otel.Tracer(name).Start(ctx, "Run")

		n, err := a.Poll(ctx)
		if err != nil {
			span.End()
			return err
		}

		a.Write(ctx, n)
		span.End()
	}
}

// Poll: asks a user for input and returns the request
func (a *App) Poll(ctx context.Context) (uint, error) {
	_, span := otel.Tracer(name).Start(ctx, "Poll")
	defer span.End()

	a.logger.Print("What Computation number would you like to know: ")

	var n uint
	_, err := fmt.Fscanf(a.reader, "%d", &n)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return 0, err
	}

	// Store
	nStr := strconv.FormatUint(uint64(n), 10)
	span.SetAttributes(attribute.String("request.n", nStr))
	return n, nil
}

// Write: writes the n-th Computation number back to the user
func (a *App) Write(ctx context.Context, n uint) {
	// span of writing data
	var spanOfWrite trace.Span
	ctx, spanOfWrite = otel.Tracer(name).Start(ctx, "Write")
	defer spanOfWrite.End()

	f, err := func(ctx context.Context) (uint64, error) {
		// span of running computation
		_, spanOfComputation := otel.Tracer(name).Start(ctx, "Computation")
		defer spanOfComputation.End()
		f, err := Computation(n)
		if err != nil {
			spanOfComputation.RecordError(err)
			spanOfComputation.SetStatus(codes.Error, err.Error())
		}
		return f, err
	}(ctx)

	if err != nil {
		a.logger.Printf("Computation(%d): %v\n", n, err)
	} else {
		a.logger.Printf("Computation(%d): %d\n", n, f)
	}
}
