package repository

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"

	"gorm.io/gorm"
)

type categoryRepoDriver struct {
	DB *gorm.DB
}

func (d *categoryRepoDriver) AddCategory(c model.Category) (category model.Category, err error) {
	category.Category = c.Category
	err = d.DB.Create(&category).Error
  return
}

func (d *categoryRepoDriver) DeleteCategory(categoryId uint) (err error) {
	err = d.DB.Delete(&model.Category{}, categoryId).Error
	return
}

func (d *categoryRepoDriver) GetCategories() (categories []model.Category, err error) {
	err = d.DB.Omit("Devices").Find(&categories).Error
	return
}

func (d *categoryRepoDriver) GetCategoryDevices(categoryId uint) (devices []model.Device, err error) {
	err = d.DB.Table("devices").Where("device_category_id = ?", categoryId).Find(&devices).Error
	return
}

func NewCategoryRepoDriver(db *gorm.DB) domain.CategoryRepository {
	return &categoryRepoDriver{
		DB: db,
	}
}
