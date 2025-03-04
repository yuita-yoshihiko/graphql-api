package converter

import (
	"graphql-api/domain/models"
	"graphql-api/domain/models/graphql"
	"graphql-api/utils"
)

type UserConverter interface {
	ConvertUserModelToGraphQLType(*models.User) (*graphql.UserDetail, error)
	ConvertUserGraphQLTypeToModel(graphql.CreateUserInput) (*models.User, error)
	ConvertUserGraphQLTypeToModelUpdate(graphql.UpdateUserInput) (*models.User, error)
	ConvertRawArgsToDBColumnNames() []string
}

type userConverterImpl struct {
}

func NewUserConverter() UserConverter {
	return &userConverterImpl{}
}

func (c *userConverterImpl) ConvertUserModelToGraphQLType(m *models.User) (*graphql.UserDetail, error) {
	return &graphql.UserDetail{
		ID:   m.ID,
		Name: m.Name,
	}, nil
}

func (c *userConverterImpl) ConvertUserGraphQLTypeToModel(input graphql.CreateUserInput) (*models.User, error) {
	return &models.User{
		Name: input.Name,
	}, nil
}

func (c *userConverterImpl) ConvertUserGraphQLTypeToModelUpdate(input graphql.UpdateUserInput) (*models.User, error) {
	return &models.User{
		Name: input.Name,
	}, nil
}

func (c userConverterImpl) ConvertRawArgsToDBColumnNames() []string {
	m := models.User{}
	i := graphql.UserDetail{}
	return utils.ConvertUpdateInputToDBColumnNames(m, i)
}
