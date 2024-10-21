package repository

import "go-belajar-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
