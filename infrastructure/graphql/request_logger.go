package graphql

import (
	"context"
	"graphql-api/constants"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
	"github.com/google/uuid"
)

func RequestLoggerHandler(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	ctx = context.WithValue(ctx, constants.RequestKey, uuid.New().String())
	oc := graphql.GetOperationContext(ctx)
	if oc.OperationName == "" {
		slog.InfoContext(ctx, "Request", "OperationName", "No match found", "UUID", ctx.Value(constants.RequestKey))
		return next(ctx)
	}
	ctx = context.WithValue(ctx, constants.OperationNameKey, oc.OperationName)
	if oc.Operation.Operation == "mutation" {
		slog.InfoContext(ctx, "Request", "OperationName", oc.OperationName, "Params", oc.Variables, "UUID", ctx.Value(constants.RequestKey))
	} else {
		slog.InfoContext(ctx, "Request", "OperationName", oc.OperationName, "UUID", ctx.Value(constants.RequestKey))
	}
	return next(ctx)
}
