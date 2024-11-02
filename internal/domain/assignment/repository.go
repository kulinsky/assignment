package assignment

import (
	"github.com/stretchr/testify/mock"
)

type Repository interface {
	Add(assignment *Assignment) error
	Find(id ID) (*Assignment, error)
}

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Add(assignment *Assignment) error {
	args := m.Called(assignment)

	return args.Error(0)
}

func (m *MockRepository) Find(id ID) (*Assignment, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).(*Assignment), args.Error(1)
}
