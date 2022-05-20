package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddCategory(t *testing.T) {
	expectedCategory := model.Category{Category: "NewCategory"}

	mock := new(domain.MockCategoryRepository)
	mock.On("AddCategory", model.Category{Category: "NewCategory"}).Return(expectedCategory, nil)

	categoryService := NewCategoryService(mock)

	_, err := categoryService.AddCategory(expectedCategory)

	assert.Nil(t, err)
	assert.Equal(t, "NewCategory", expectedCategory.Category)
	mock.AssertNumberOfCalls(t, "AddCategory", 1)
}

func TestDeleteCategory(t *testing.T) {
	inputCategoryId := uint(1)

	mock := new(domain.MockCategoryRepository)
	mock.On("DeleteCategory", inputCategoryId).Return(nil)

	categoryService := NewCategoryService(mock)

	err := categoryService.DeleteCategory(uint(inputCategoryId))

	assert.Nil(t, err)
}

func TestGetCategories(t *testing.T) {
	expectedCategories := []model.Category{
		{
			Category: "Category1",
		},
		{
			Category: "Category2",
		},
		{
			Category: "Category3",
		},
	}

	mock := new(domain.MockCategoryRepository)
	mock.On("GetCategories").Return(expectedCategories, nil)

	categoryService := NewCategoryService(mock)

	categories, err := categoryService.GetCategories()

	assert.Nil(t, err)
	assert.Equal(t, 3, len(categories))
}

func TestGetCategoryDevices(t *testing.T) {
	expectedDevices := []model.Device{
		{
			CategoryID: 1,
		},
		{
			CategoryID: 1,
		},
	}
	inputArg := uint(1)

	mock := new(domain.MockCategoryRepository)
	mock.On("GetCategoryDevices", inputArg).Return(expectedDevices, nil)

	categoryService := NewCategoryService(mock)

	devices, err := categoryService.GetCategoryDevices(inputArg)

	assert.Nil(t, err)
	assert.NotNil(t, devices[0].CategoryID)
	assert.Equal(t, uint(1), devices[0].CategoryID)
}
