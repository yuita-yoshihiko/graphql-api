package converter

import (
	"graphql-api/domain/models"
	"graphql-api/domain/models/graphql"
	"graphql-api/utils"
)

type PostConverter interface {
	ConvertPostModelToGraphQLType(*models.Post) (*graphql.PostDetail, error)
	ConvertPostModelsToGraphQLTypes([]*models.Post) ([]*graphql.PostDetail, error)
	ConvertPostGraphQLTypeToModel(graphql.CreatePostInput) (*models.Post, error)
	ConvertPostGraphQLTypeToModelUpdate(graphql.UpdatePostInput) (*models.Post, error)
	ConvertRawArgsToDBColumnNames() []string
}

type postConverterImpl struct {
}

func NewPostConverter() PostConverter {
	return &postConverterImpl{}
}

func (c *postConverterImpl) ConvertPostModelToGraphQLType(m *models.Post) (*graphql.PostDetail, error) {
	return &graphql.PostDetail{
		ID: m.ID,
		User: &graphql.UserDetail{
			ID: m.UserID,
		},
		Title:   m.Title,
		Content: m.Content,
	}, nil
}

func (c *postConverterImpl) ConvertPostModelsToGraphQLTypes(ms []*models.Post) ([]*graphql.PostDetail, error) {
	postDetails := []*graphql.PostDetail{}
	for _, m := range ms {
		postDetail, err := c.ConvertPostModelToGraphQLType(m)
		if err != nil {
			return nil, err
		}
		postDetails = append(postDetails, postDetail)
	}
	return postDetails, nil
}

func (c *postConverterImpl) ConvertPostGraphQLTypeToModel(input graphql.CreatePostInput) (*models.Post, error) {
	return &models.Post{
		UserID:  input.UserID,
		Title:   input.Title,
		Content: input.Content,
	}, nil
}

func (c *postConverterImpl) ConvertPostGraphQLTypeToModelUpdate(input graphql.UpdatePostInput) (*models.Post, error) {
	return &models.Post{
		Title:   input.Title,
		Content: input.Content,
	}, nil
}

func (c postConverterImpl) ConvertRawArgsToDBColumnNames() []string {
	m := models.Post{}
	i := graphql.PostDetail{}
	return utils.ConvertUpdateInputToDBColumnNames(m, i)
}
