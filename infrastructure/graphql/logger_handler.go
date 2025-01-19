package graphql

import (
	"context"
	"graphql-api/constants"
	"log/slog"

	"github.com/99designs/gqlgen/graphql"
)

func LoggerHandler(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
	response := next(ctx)
	if response.Errors == nil {
		slog.InfoContext(ctx, "GraphQL Operation", "OperationName", response.Data, "UUID", ctx.Value(constants.RequestKey))
	} else {
		slog.InfoContext(ctx, "GraphQL Operation", "Error", response.Errors, "UUID", ctx.Value(constants.RequestKey))
	}
	return response
}
