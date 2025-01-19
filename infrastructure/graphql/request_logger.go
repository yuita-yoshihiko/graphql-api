package graphql

import (
	"context"
	"graphql-api/constants"
	"log/slog"
	"regexp"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

var re = regexp.MustCompile(`(?m)^\s*(?:query|mutation|subscription)?\s*{\s*(\w+)`)

func RequestLoggerHandler(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)
	ctx = context.WithValue(ctx, constants.RequestKey, uuid.New().String())
	matches := re.FindStringSubmatch(oc.RawQuery)
	if len(matches) > 1 {
		slog.InfoContext(ctx, "Request", "OperationName", matches[1], "UUID", ctx.Value(constants.RequestKey))
	} else {
		slog.InfoContext(ctx, "Request", "OperationName", "No match found", "UUID", ctx.Value(constants.RequestKey))
	}
	ctx = context.WithValue(ctx, constants.OperationNameKey, matches[1])
	return next(ctx)
}
