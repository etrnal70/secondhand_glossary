package rest

import (
	"net/http"
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/domain"
	j "secondhand_glossary/internal/middleware/jwt"
	"secondhand_glossary/internal/model"
	"strconv"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type DeviceController struct {
	conf config.Config
	s    domain.DeviceService
}

// AddDevice godoc
// @Summary Add new device
// @Tags device
// @Accept json
// @Produce json
// @Success 200 {object} model.Device
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device [post]
func (c *DeviceController) AddDeviceController(ctx echo.Context) error {
	device := model.Device{}
	err := ctx.Bind(&device)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newDevice, err := c.s.AddDevice(device)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting data : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"device":  newDevice,
	})
}

// AddDeviceLink godoc
// @Summary Add new device's link
// @Tags device
// @Accept json
// @Produce json
// @Param id path uint true "Device ID"
// @Success 200 {object} model.Link
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/link [post]
func (c *DeviceController) AddDeviceLinkController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	link := model.Link{DeviceID: uint(deviceId)}
	err = ctx.Bind(&link)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newLink, err := c.s.AddDeviceLink(link)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting data : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"link":    newLink,
	})
}

// AddDeviceReview godoc
// @Summary Add new device's review
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Success 200 {object} model.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/review [post]
func (c *DeviceController) AddDeviceReviewController(ctx echo.Context) error {
	// TODO Need user id from token
	user := ctx.Get("user").(*jwt.Token)
	userClaims := user.Claims.(j.CustomClaims)
	userId := userClaims.UserID

	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	review := model.Review{UserID: userId, DeviceID: uint(deviceId)}
	err = ctx.Bind(&review)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newReview, err := c.s.AddDeviceReview(review)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error inserting review : " + err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"review":  newReview,
	})
}

// AddDeviceTrait godoc
// @Summary Add device's trait
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param trait_id path uint true "DeviceTrait ID"
// @Success 200 {object} model.Device
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/trait/{trait_id} [post]
func (c *DeviceController) AddDeviceTraitController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	traitIdStr := ctx.Param("trait_id")
	traitId, err := strconv.Atoi(traitIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	newDevice, err := c.s.AddDeviceTrait(uint(deviceId), uint(traitId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error adding device's trait : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"device":  newDevice,
	})
}

// DeleteDevice godoc
// @Summary Delete device by id
// @Tags device
// @Produce json
// @Param id path uint true "Device ID"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{id} [delete]
func (c *DeviceController) DeleteDeviceController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	err = c.s.DeleteDevice(uint(deviceId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting device : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success deleting device",
	})
}

// DeleteDeviceLink godoc
// @Summary Delete device link by id
// @Tags device
// @Produce json
// @Param id path uint true "Device ID"
// @Success 200 {object} model.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/link/{link_id} [delete]
func (c *DeviceController) DeleteDeviceLinkController(ctx echo.Context) error {
	// TODO Hadehh
	// deviceIdStr := ctx.Param("device_id")
	// deviceId, err := strconv.Atoi(deviceIdStr)
	// if err != nil {
	// log.Error(err)
	// 	return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
	// 		"message": "Error on request parameter : " + err.Error(),
	// 	})
	// }

	linkIdStr := ctx.Param("trait_id")
	linkId, err := strconv.Atoi(linkIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	err = c.s.DeleteDeviceLink(uint(linkId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting device's link : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success deleting device's link",
	})
}

// DeleteDeviceReview godoc
// @Summary Delete device's review by id
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param review_id path uint true "Device's Review ID"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/review/{review_id} [delete]
func (c *DeviceController) DeleteDeviceReviewController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	reviewIdStr := ctx.Param("review_id")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	err = c.s.DeleteDeviceReview(model.Review{ID: uint(reviewId), DeviceID: uint(deviceId)})
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting device's review : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success deleting device's review",
	})
}

// DeleteDeviceTrait godoc
// @Summary Delete device's trait by id
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param trait_id path uint true "Device' Trait ID"
// @Success 200
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/trait/{trait_id} [delete]
func (c *DeviceController) DeleteDeviceTraitController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	traitIdStr := ctx.Param("trait_id")
	traitId, err := strconv.Atoi(traitIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	err = c.s.DeleteDeviceTrait(uint(deviceId), uint(traitId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error removing device from trait : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success removing device from trait",
	})
}

// EditDevice godoc
// @Summary Edit device by id
// @Tags device
// @Accept json
// @Produce json
// @Param id path uint true "Device ID"
// @Success 200 {object} model.Device
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{id} [put]
func (c *DeviceController) EditDeviceController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	device := model.Device{ID: uint(deviceId)}
	err = ctx.Bind(&device)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newDevice, err := c.s.EditDevice(device)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating device : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"device":  newDevice,
	})
}

// EditDeviceLink godoc
// @Summary Add new device's review
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param link_id path uint true "Device's Link ID"
// @Success 200 {object} model.Link
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/link/{link_id} [put]
func (c *DeviceController) EditDeviceLinkController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	linkIdStr := ctx.Param("link_id")
	linkId, err := strconv.Atoi(linkIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	link := model.Link{}
	err = ctx.Bind(&link)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newLink, err := c.s.EditDeviceLink(model.Link{ID: uint(linkId), DeviceID: uint(deviceId)})
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating device's link : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"device":  newLink,
	})
}

// EditDeviceReview godoc
// @Summary Edit device's review by id
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param review_id path uint true "Device's Review ID"
// @Success 200 {object} model.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/review/{review_id} [put]
func (c *DeviceController) EditDeviceReviewController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	reviewIdStr := ctx.Param("review_id")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	review := model.Review{ID: uint(reviewId), DeviceID: uint(deviceId)}
	err = ctx.Bind(&review)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newReview, err := c.s.EditDeviceReview(review)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating device's review : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"review":  newReview,
	})
}

// GetDevice godoc
// @Summary Get device by id
// @Tags device
// @Produce json
// @Param id path uint true "Device ID"
// @Success 200 {object} model.Device
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id} [get]
func (c *DeviceController) GetDeviceController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	device, err := c.s.GetDevice(uint(deviceId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"device":  device,
	})
}

// GetDeviceLink godoc
// @Summary Get device's link by id
// @Tags device
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param link_id path uint true "Device's Link ID"
// @Success 200 {object} model.Link
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/link/{link_id} [get]
func (c *DeviceController) GetDeviceLinkController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	linkIdStr := ctx.Param("link_id")
	linkId, err := strconv.Atoi(linkIdStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	link, err := c.s.GetDeviceLink(uint(deviceId), uint(linkId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device's link : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"link":    link,
	})
}

// GetDeviceLinks godoc
// @Summary Get all device's links
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Success 200 {object} []model.Link
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/link [get]
func (c *DeviceController) GetDeviceLinksController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	links, err := c.s.GetDeviceLinks(uint(deviceId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device's links : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"link":    links,
	})
}

// GetDeviceReview godoc
// @Summary Get device's review by id
// @Tags device
// @Produce json
// @Param device_id path uint true "Device ID"
// @Param review_id path uint true "Device's Review ID"
// @Success 200 {object} model.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/review/{review_id} [get]
func (c *DeviceController) GetDeviceReviewController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	reviewIdStr := ctx.Param("review_id")
	reviewId, err := strconv.Atoi(reviewIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	review, err := c.s.GetDeviceReview(uint(deviceId), uint(reviewId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device's review : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"review":  review,
	})
}

// GetDeviceReviews godoc
// @Summary Get all device's reviews
// @Tags device
// @Produce json
// @Param device_id path uint true "Device ID"
// @Success 200 {object} []model.Review
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/review [get]
func (c *DeviceController) GetDeviceReviewsController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	reviews, err := c.s.GetDeviceReviews(uint(deviceId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device's reviews : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"reviews": reviews,
	})
}

// GetDeviceTraits godoc
// @Summary Get all device's traits
// @Tags device
// @Produce json
// @Param device_id path uint true "Device ID"
// @Success 200 {object} []model.Trait
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{id}/trait [get]
func (c *DeviceController) GetDeviceTraitsController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	traits, err := c.s.GetDeviceTraits(uint(deviceId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device's traits : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"traits":  traits,
	})
}

// GetDevices godoc
// @Summary Get all device details
// @Tags device
// @Produce json
// @Success 200 {object} []model.Device
// @Failure 500 {object} map[string]interface{}
// @Router /device [get]
func (c *DeviceController) GetDevicesController(ctx echo.Context) error {
	devices, err := c.s.GetDevices()
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting devices : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"devices": devices,
	})
}

// EditDeviceScore godoc
// @Summary Update device's score
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Success 200 {object} model.Scores
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/score [put]
func (c *DeviceController) EditDeviceScoreController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	score := model.Scores{DeviceID: uint(deviceId)}
	err = ctx.Bind(&score)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request :" + err.Error(),
		})
	}

	newScore, err := c.s.EditDeviceScore(score)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating device's score : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"score":   newScore,
	})
}

// GetDeviceScore godoc
// @Summary Update device's score
// @Tags device
// @Accept json
// @Produce json
// @Param device_id path uint true "Device ID"
// @Success 200 {object} model.Scores
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /device/{device_id}/score [get]
func (c *DeviceController) GetDeviceScoreController(ctx echo.Context) error {
	deviceIdStr := ctx.Param("device_id")
	deviceId, err := strconv.Atoi(deviceIdStr)
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Error on request parameter : " + err.Error(),
		})
	}

	score, err := c.s.GetDeviceScore(uint(deviceId))
	if err != nil {
		log.Error(err)
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error getting device's score : " + err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"score":   score,
	})
}
