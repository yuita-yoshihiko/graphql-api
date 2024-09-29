package dataloader

import (
	"context"
	"time"

	"github.com/labstack/echo/v4"
	"graphql-api/infrastructure/db"
	"graphql-api/interface/database"
	"graphql-api/interface/database/dataloader"
	"graphql-api/usecase/converter"
)

// DataLoaderMiddleWare dataloaderのcontextへの登録
func DataLoaderMiddleWare() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			dbAdministrator := db.NewDBAdministrator(db.DB)
			ctx := context.WithValue(
				c.Request().Context(),
				dataloader.CDataloaderKey,
				dataloader.NewCommentDataLoader(
					converter.NewCommentConverter(),
					database.NewCommentRepository(dbAdministrator),
					100*time.Millisecond,
				).NewInterface(),
			)
			c.SetRequest(c.Request().WithContext(ctx))
			return next(c)
		}
	}
}
