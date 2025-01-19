package main

import (
	"graphql-api/infrastructure/db"
	"graphql-api/interface/resolvers"
	"graphql-api/internal/middleware/dataloader"
	"graphql-api/route"
	"log"
	"log/slog"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.Init()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	e.GET("/", route.PlaygroundHandler())
	g := e.Group("/query")

	g.Use(dataloader.DataLoaderMiddleWare())

	resolver := resolvers.NewResolver()

	g.POST("/", route.DefineGraphQL(resolver))

	log.Fatal(e.Start(":8080"))
}
