package converter

import (
	"graphql-api/domain/models"
	"graphql-api/domain/models/graphql"
)

type UserConverter interface {
	ConvertUserModelToGraphQLType(*models.User) (*graphql.UserDetail, error)
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
