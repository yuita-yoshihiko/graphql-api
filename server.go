package main

import (
	"os"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gin-gonic/gin"
)

const defaultPort = "8888"

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	r := gin.Default()

	r.GET("/", playgroundHandler())
	r.Run(":" + port)
}
