package repository

import (
	"github.com/stretchr/testify/mock"
	"go-belajar-unit-test/entity"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (repository *CategoryRepositoryMock) FindById(id string) *entity.Category {
	arguments := repository.Mock.Called(id) // mengembalikan value
	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}
}
