package resolvers

import (
	"graphql-api/infrastructure/db"
	"graphql-api/interface/resolvers/database"
	"graphql-api/usecase"
	"graphql-api/usecase/converter"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	UserUsecase usecase.UserUsecase
}

func NewResolver() *Resolver {
	dbAdministrator := db.NewDBAdministrator(db.DB)
	userRepository := database.NewUserRepository(dbAdministrator)
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		converter.NewUserConverter(),
	)
	return &Resolver{
		UserUsecase: userUsecase,
	}
}
