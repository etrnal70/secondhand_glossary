package rest

import (
	"net/http"
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type CategoryController struct {
	conf config.Config
	s    domain.CategoryService
}

// AddCategory godoc
// @Summary Add new category
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} model.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /category [post]
func (c *CategoryController) AddCategoryController(ctx echo.Context) error {
	category := model.Category{}
	err := ctx.Bind(&category)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newCategory, err := c.s.AddCategory(category)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error adding category : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, newCategory)
}

// DeleteCategory godoc
// @Summary Delete category by id
// @Tags category
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /category [delete]
func (c *CategoryController) DeleteCategoryController(ctx echo.Context) error {
	categoryIdStr := ctx.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}
	err = c.s.DeleteCategory(uint(categoryId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting category : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success deleting category",
	})
}

// GetCategories godoc
// @Summary Get all categories
// @Tags category
// @Produce json
// @Success 200 {object} []model.Category
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /category [get]
func (c *CategoryController) GetCategoriesController(ctx echo.Context) error {
	categories, err := c.s.GetCategories()
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting categories : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, categories)
}

// GetCategoryDevices godoc
// @Summary Add new device
// @Tags category
// @Accept json
// @Produce json
// @Success 200 {object} []model.Device
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /category/{category_id} [get]
func (c *CategoryController) GetCategoryDevicesController(ctx echo.Context) error {
	categoryIdStr := ctx.Param("category_id")
	categoryId, err := strconv.Atoi(categoryIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	devices, err := c.s.GetCategoryDevices(uint(categoryId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting category devices : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, devices)
}
