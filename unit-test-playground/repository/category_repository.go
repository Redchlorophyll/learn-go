package repository

import "unit-test-playground/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
