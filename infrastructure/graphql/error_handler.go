package graphql

import (
	"context"
	"log/slog"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type ErrorPresenterFunc func(ctx context.Context, err error) *gqlerror.Error
type ExtendedError interface {
	Extensions() map[string]interface{}
}

func ErrorHandler(ctx context.Context, err error) *gqlerror.Error {
	if gqlerr, ok := err.(*gqlerror.Error); ok {
		if gqlerr.Path == nil {
			gqlerr.Path = graphql.GetFieldContext(ctx).Path()
		}
		logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))
		slog.SetDefault(logger)
		slog.InfoContext(ctx, "errorMessage", "err", gqlerr.Message, "path", gqlerr.Path)
		return gqlerr
	}

	var extensions map[string]interface{}
	if ee, ok := err.(ExtendedError); ok {
		extensions = ee.Extensions()
	}

	return &gqlerror.Error{
		Message:    err.Error(),
		Path:       graphql.GetFieldContext(ctx).Path(),
		Extensions: extensions,
	}
}
