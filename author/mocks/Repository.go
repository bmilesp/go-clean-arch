// Code generated by mockery v1.0.0
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import models "github.com/bmilesp/go-clean-arch/models"

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByID provides a mock function with given fields: ctx, id
func (_m *Repository) GetByID(ctx context.Context, id int64) (*models.Author, error) {
	ret := _m.Called(ctx, id)

	var r0 *models.Author
	if rf, ok := ret.Get(0).(func(context.Context, int64) *models.Author); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Author)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
