package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
)

type categoryService struct {
	Repo domain.CategoryRepository
}

func (s *categoryService) AddCategory(c model.Category) (category model.Category, err error) {
  category, err = s.Repo.AddCategory(c)
  return
}

func (s *categoryService) DeleteCategory(categoryId uint) (err error) {
  err = s.Repo.DeleteCategory(categoryId)
  return
}

func (s *categoryService) GetCategories() (categories []model.Category, err error) {
	categories, err = s.Repo.GetCategories()
	return
}

func (s *categoryService) GetCategoryDevices(categoryId uint) (devices []model.Device, err error) {
	devices, err = s.Repo.GetCategoryDevices(categoryId)
	return
}

func NewCategoryService(r domain.CategoryRepository) domain.CategoryService {
	return &categoryService{
		Repo: r,
	}
}
