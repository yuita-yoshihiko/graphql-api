package converter

import (
	"graphql-api/domain/models"
	"graphql-api/domain/models/graphql"
	"graphql-api/utils"
)

type StaffConverter interface {
	ToStaffModelFromCreateInput(graphql.StaffCreateInput) (*models.Staff, error)
	ToStaffModelFromUpdateInput(graphql.StaffUpdateInput) (*models.Staff, error)
	ToStaffDetail(*models.Staff) (*graphql.StaffDetail, error)
	ToStaffDetails([]*models.Staff) ([]*graphql.StaffDetail, error)
	ToDBColumnsFromGraphQLFields() []string
}

type staffConverterImpl struct {
}

// NewStaffConverter StaffConverterを生成
func NewStaffConverter() StaffConverter {
	return &staffConverterImpl{}
}

func (c staffConverterImpl) ToStaffModelFromCreateInput(input graphql.StaffCreateInput) (*models.Staff, error) {
	return nil, nil
}

func (c staffConverterImpl) ToStaffModelFromUpdateInput(input graphql.StaffUpdateInput) (*models.Staff, error) {
	return nil, nil
}

func (c staffConverterImpl) ToStaffDetail(m *models.Staff) (*graphql.StaffDetail, error) {
	if m == nil {
		return nil, nil
	}
	return &graphql.StaffDetail{}, nil
}

func (c staffConverterImpl) ToStaffDetails(ms []*models.Staff) ([]*graphql.StaffDetail, error) {
	details := []*graphql.StaffDetail{}
	for _, m := range ms {
		detail, err := c.ToStaffDetail(m)
		if err != nil {
			return nil, err
		} else if detail == nil {
			continue
		}
		details = append(details, detail)
	}
	return details, nil
}

func (c staffConverterImpl) ToDBColumnsFromGraphQLFields() []string {
	m := models.Staff{}
	i := graphql.StaffDetail{}
	return utils.ConvertUpdateInputToDBColumnNames(m, i)
}
