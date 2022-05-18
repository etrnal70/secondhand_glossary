package rest

import (
	"net/http"
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TraitController struct {
	conf config.Config
	s    domain.TraitService
}

// AddTrait godoc
// @Summary Add new trait
// @Tags trait
// @Accept json
// @Produce json
// @Success 200 {object} model.Trait
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /trait [post]
func (c *TraitController) AddTraitController(ctx echo.Context) error {
	trait := model.Trait{}
	err := ctx.Bind(&trait)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newTrait, err := c.s.AddTrait(trait)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error adding new trait : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success adding new trait",
		"trait":   newTrait,
	})
}

// EditTrait godoc
// @Summary Edit trait by id
// @Tags trait
// @Accept json
// @Produce json
// @Success 200 {object} model.Trait
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /trait/{trait_id} [put]
func (c *TraitController) EditTraitController(ctx echo.Context) error {
	traitIdStr := ctx.Param("trait_id")
	traitId, err := strconv.Atoi(traitIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	trait := model.Trait{ID: uint(traitId)}
	err = ctx.Bind(&trait)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newTrait, err := c.s.EditTrait(trait)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating trait : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success updating trait",
		"trait":   newTrait,
	})
}

// GetTraits godoc
// @Summary Get all traits
// @Tags trait
// @Produce json
// @Success 200 {object} []model.Trait
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /trait [get]
func (c *TraitController) GetTraitsController(ctx echo.Context) error {
	traits, err := c.s.GetTraits()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting traits : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"traits":  traits,
	})
}

// GetTraitDevices godoc
// @Summary Get all trait devices
// @Tags trait
// @Produce json
// @Success 200 {object} []model.Device
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /trait/{trait_id} [get]
func (c *TraitController) GetTraitDevicesController(ctx echo.Context) error {
	traitIdStr := ctx.Param("trait_id")
	traitId, err := strconv.Atoi(traitIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	devices, err := c.s.GetTraitDevices(uint(traitId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting trait's devices : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"devices": devices,
	})
}

// DeleteTrait godoc
// @Summary Delete trait by id
// @Tags trait
// @Produce json
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /trait/{trait_id} [delete]
func (c *TraitController) DeleteTraitController(ctx echo.Context) error {
	traitIdStr := ctx.Param("trait_id")
	traitId, err := strconv.Atoi(traitIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	err = c.s.DeleteTrait(uint(traitId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting trait : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success deleting trait",
	})
}
