package domain

import "secondhand_glossary/internal/model"

type DeviceRepository interface {
	AddDevice(d model.Device) (device model.Device, err error)
	GetDevice(deviceId uint) (device model.Device, err error)
	GetDevices() (devices []model.Device, err error)
	EditDevice(d model.Device) (device model.Device, err error)
	DeleteDevice(deviceId uint) (err error)

	AddDeviceReview(r model.Review) (review model.Review, err error)
	GetDeviceReview(deviceId uint, reviewId uint) (review model.Review, err error)
	GetDeviceReviews(deviceId uint) (reviews []model.Review, err error)
	EditDeviceReview(r model.Review) (review model.Review, err error)
	DeleteDeviceReview(r model.Review) (err error)

	AddDeviceTrait(deviceId uint, traitId uint) (device model.Device, err error)
	GetDeviceTraits(deviceId uint) (traits []*model.Trait, err error)
	DeleteDeviceTrait(deviceId uint, traitId uint) (err error)

	AddDeviceLink(b model.Link) (link model.Link, err error)
	GetDeviceLink(deviceId uint, linkId uint) (link model.Link, err error)
	GetDeviceLinks(deviceId uint) (links []*model.Link, err error)
	EditDeviceLink(b model.Link) (link model.Link, err error)
	DeleteDeviceLink(linkId uint) (err error)

	// Category

	// Score (External service only)
	GetDeviceScore(deviceId uint) (score model.Scores, err error)
	EditDeviceScore(scr model.Scores) (score model.Scores, err error)
}

type DeviceService interface {
	AddDevice(d model.Device) (device model.Device, err error)
	GetDevice(deviceId uint) (device model.Device, err error)
	GetDevices() (devices []model.Device, err error)
	EditDevice(d model.Device) (device model.Device, err error)
	DeleteDevice(deviceId uint) (err error)

	AddDeviceReview(r model.Review) (review model.Review, err error)
	GetDeviceReview(deviceId uint, reviewId uint) (review model.Review, err error)
	GetDeviceReviews(deviceId uint) (reviews []model.Review, err error)
	EditDeviceReview(r model.Review) (review model.Review, err error)
	DeleteDeviceReview(r model.Review) (err error)

	AddDeviceTrait(deviceId uint, traitId uint) (device model.Device, err error)
	GetDeviceTraits(deviceId uint) (traits []*model.Trait, err error)
	DeleteDeviceTrait(deviceId uint, traitId uint) (err error)

	AddDeviceLink(b model.Link) (link model.Link, err error)
	GetDeviceLink(deviceId uint, linkId uint) (link model.Link, err error)
	GetDeviceLinks(deviceId uint) (links []*model.Link, err error)
	EditDeviceLink(b model.Link) (link model.Link, err error)
	DeleteDeviceLink(linkId uint) (err error)

	// Category

	// Scores (External service only)
	GetDeviceScore(deviceId uint) (score model.Scores, err error)
	EditDeviceScore(scr model.Scores) (score model.Scores, err error)
}
