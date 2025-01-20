package graphql

import (
	"context"
	"graphql-api/constants"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func RequestLoggerHandler(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	oc := graphql.GetOperationContext(ctx)
	ctx = context.WithValue(ctx, constants.RequestKey, uuid.New().String())
	if oc.OperationName != "" {
		slog.InfoContext(ctx, "Request", "OperationName", oc.OperationName, "UUID", ctx.Value(constants.RequestKey))
	} else {
		slog.InfoContext(ctx, "Request", "OperationName", "No match found", "UUID", ctx.Value(constants.RequestKey))
	}
	ctx = context.WithValue(ctx, constants.OperationNameKey, oc.OperationName)
	return next(ctx)
}
