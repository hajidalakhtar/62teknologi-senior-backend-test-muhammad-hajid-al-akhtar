package mocks

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BusinessRepository struct {
	Mock mock.Mock
}

func NewBusinessRepository(mock mock.Mock) domain.BusinessRepository {
	return &BusinessRepository{Mock: mock}
}

func (_m BusinessRepository) Find(ctx context.Context, term string, sortBy string, limit int, offset int, latitude float64, longitude float64) ([]domain.Business, error) {
	ret := _m.Mock.Called(ctx, term, sortBy, limit, offset, latitude, longitude)

	var r0 []domain.Business
	var r1 error

	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int, float64, float64) ([]domain.Business, error)); ok {
		r0, r1 = rf(ctx, term, sortBy, limit, offset, latitude, longitude)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Business)
		}
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (_m BusinessRepository) Update(ctx context.Context, bs *domain.Business, id uuid.UUID) error {
	ret := _m.Mock.Called(ctx, bs, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Business, uuid.UUID) error); ok {
		r0 = rf(ctx, bs, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m BusinessRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Mock.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
