package resolvers

import (
	"graphql-api/infrastructure/db"
	"graphql-api/interface/database"
	"graphql-api/usecase"
	"graphql-api/usecase/converter"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UserUsecase    usecase.UserUsecase
	PostUsecase    usecase.PostUsecase
	CommentUsecase usecase.CommentUsecase
}

func NewResolver() *Resolver {
	dbAdministrator := db.NewDBAdministrator(db.DB)
	userRepository := database.NewUserRepository(dbAdministrator)
	userUsecase := usecase.NewUserUsecase(
		userRepository,
		converter.NewUserConverter(),
	)
	postRepository := database.NewPostRepository(dbAdministrator)
	postUsecase := usecase.NewPostUsecase(
		postRepository,
		converter.NewPostConverter(),
	)
	commentRepository := database.NewCommentRepository(dbAdministrator)
	commentUsecase := usecase.NewCommentUsecase(
		commentRepository,
		converter.NewCommentConverter(),
	)
	return &Resolver{
		UserUsecase:    userUsecase,
		PostUsecase:    postUsecase,
		CommentUsecase: commentUsecase,
	}
}
