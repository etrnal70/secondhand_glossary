package domain

import "secondhand_glossary/internal/model"

type CategoryRepository interface {
	AddCategory(c model.Category) (category model.Category, err error)
	GetCategories() (categories []model.Category, err error)
	GetCategoryDevices(categoryId uint) (devices []model.Device, err error)
  DeleteCategory(categoryId uint) (err error)
}

type CategoryService interface {
	AddCategory(c model.Category) (category model.Category, err error)
	GetCategories() (categories []model.Category, err error)
	GetCategoryDevices(categoryId uint) (devices []model.Device, err error)
  DeleteCategory(categoryId uint) (err error)
}
