package graphql

import (
	"context"
	"graphql-api/constants"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
)

func ResponseLoggerHandler(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	response := next(ctx)
	if response.Errors == nil {
		slog.InfoContext(ctx, "GraphQL Operation", "OperationName", ctx.Value(constants.OperationNameKey), "UUID", ctx.Value(constants.RequestKey))
	} else {
		slog.ErrorContext(ctx, "GraphQL Operation", "OperationName", ctx.Value(constants.OperationNameKey), "ErrorInfo", response.Errors, "UUID", ctx.Value(constants.RequestKey))
	}
	return response
}
