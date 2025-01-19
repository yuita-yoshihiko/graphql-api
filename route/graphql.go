package route

import (
	"graphql-api/infrastructure/graphql"
	"graphql-api/interface/resolvers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func DefineGraphQL(r *resolvers.Resolver) echo.HandlerFunc {
	config := graphql.Config{
		Resolvers: r,
	}
	handler := handler.NewDefaultServer(graphql.NewExecutableSchema(config))
	handler.AroundResponses(graphql.LoggerHandler)
	return func(c echo.Context) error {
		handler.ServeHTTP(c.Response().Writer, c.Request())
		return nil
	}
}

func PlaygroundHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		playground.Handler("GraphQL", "/query/").ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
