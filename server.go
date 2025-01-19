package main

import (
	"graphql-api/infrastructure/db"
	"graphql-api/interface/resolvers"
	"graphql-api/internal/middleware/dataloader"
	"graphql-api/internal/middleware/logger"
	"graphql-api/route"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	db.Init()
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", route.PlaygroundHandler())
	g := e.Group("/query")

	g.Use(dataloader.DataLoaderMiddleWare())
	g.Use(logger.LoggerMiddleWare())

	resolver := resolvers.NewResolver()

	g.POST("/", route.DefineGraphQL(resolver))

	log.Fatal(e.Start(":8080"))
}
