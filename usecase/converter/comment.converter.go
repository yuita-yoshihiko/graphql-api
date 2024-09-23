package converter

import (
	"graphql-api/domain/models"
	"graphql-api/domain/models/graphql"
	"graphql-api/utils"
)

type CommentConverter interface {
	ConvertCommentModelToGraphQLType(*models.Comment) (*graphql.CommentDetail, error)
	ConvertCommentModelsToGraphQLTypes([]*models.Comment) ([]*graphql.CommentDetail, error)
	ConvertCommentGraphQLTypeToModel(graphql.CreateCommentInput) (*models.Comment, error)
	ConvertCommentGraphQLTypeToModelUpdate(graphql.UpdateCommentInput) (*models.Comment, error)
	ConvertRawArgsToDBColumnNames(map[string]interface{}) ([]string, error)
}

type commentConverterImpl struct {
}

func NewCommentConverter() CommentConverter {
	return &commentConverterImpl{}
}

func (c *commentConverterImpl) ConvertCommentModelToGraphQLType(m *models.Comment) (*graphql.CommentDetail, error) {
	return &graphql.CommentDetail{
		ID:   m.ID,
		User: &graphql.UserDetail{
			ID: m.UserID,
		},
		Post: &graphql.PostDetail{
			ID: m.PostID,
		},
		Content: m.Content,
	}, nil
}

func (c *commentConverterImpl) ConvertCommentModelsToGraphQLTypes(ms []*models.Comment) ([]*graphql.CommentDetail, error) {
	commentDetails := []*graphql.CommentDetail{}
	for _, m := range ms {
		commentDetail, err := c.ConvertCommentModelToGraphQLType(m)
		if err != nil {
			return nil, err
		}
		commentDetails = append(commentDetails, commentDetail)
	}
	return commentDetails, nil
}

func (c *commentConverterImpl) ConvertCommentGraphQLTypeToModel(input graphql.CreateCommentInput) (*models.Comment, error) {
	return &models.Comment{
		UserID: input.UserID,
		PostID: input.PostID,
		Content: input.Content,
	}, nil
}

func (c *commentConverterImpl) ConvertCommentGraphQLTypeToModelUpdate(input graphql.UpdateCommentInput) (*models.Comment, error) {
	return &models.Comment{
		Content: input.Content,
	}, nil
}

func (c commentConverterImpl) ConvertRawArgsToDBColumnNames(rawArgs map[string]interface{}) ([]string, error) {
	i := graphql.CommentDetail{}
	m := models.Comment{}
	return utils.ConvertRawArgsToColumnNames(rawArgs, i, m)
}
