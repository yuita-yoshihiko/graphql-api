package usecase

import (
	"context"
	"graphql-api/domain/models/graphql"
	"graphql-api/usecase/converter"
	"graphql-api/usecase/repository"
)

type StaffUseCase interface {
	Fetch(context.Context, int64) (*graphql.StaffDetail, error)
	Create(context.Context, graphql.StaffCreateInput) (*graphql.StaffDetail, error)
	Update(context.Context, graphql.StaffUpdateInput) (*graphql.StaffDetail, error)
}

type staffUseCaseImpl struct {
	repository repository.StaffRepository
	converter  converter.StaffConverter
}

func NewStaffUseCase(r repository.StaffRepository,
	c converter.StaffConverter,
) StaffUseCase {
	return &staffUseCaseImpl{
		repository: r,
		converter:  c,
	}
}

func (u *staffUseCaseImpl) Fetch(ctx context.Context, id int64) (*graphql.StaffDetail, error) {
	m, err := u.repository.Fetch(ctx, id)
	if err != nil {
		return nil, err
	}
	return u.converter.ToStaffDetail(m)
}

func (u *staffUseCaseImpl) Create(ctx context.Context, input graphql.StaffCreateInput) (*graphql.StaffDetail, error) {
	m, err := u.converter.ToStaffModelFromCreateInput(input)
	if err != nil {
		return nil, err
	}
	if err := u.repository.Create(ctx, m); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
}

func (u *staffUseCaseImpl) Update(ctx context.Context, input graphql.StaffUpdateInput) (*graphql.StaffDetail, error) {
	columns := u.converter.ToDBColumnsFromGraphQLFields()
	m, err := u.converter.ToStaffModelFromUpdateInput(input)
	if err != nil {
		return nil, err
	}
	m.ID = input.ID
	if err := u.repository.Update(ctx, m, columns); err != nil {
		return nil, err
	}
	return u.Fetch(ctx, m.ID)
}
