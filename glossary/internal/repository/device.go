package repository

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type deviceRepoDriver struct {
	DB *gorm.DB
}

func (d *deviceRepoDriver) AddDevice(dev model.Device) (device model.Device, err error) {
	device = dev
	device.Scores = model.Scores{
		UserScore:    0,
		CrawlerScore: 0,
	}
	err = d.DB.Preload(clause.Associations).Omit("Scores", "Category", "Reviews", "Links").Create(&device).Error
	return
}

func (d *deviceRepoDriver) DeleteDevice(deviceId uint) (err error) {
	err = d.DB.Delete(&model.Device{}, deviceId).Error
	return
}

func (d *deviceRepoDriver) EditDevice(dev model.Device) (device model.Device, err error) {
	device = dev
	err = d.DB.Save(&device).Error
	return
}

func (d *deviceRepoDriver) AddDeviceLink(b model.Link) (link model.Link, err error) {
	link = b
	err = d.DB.Save(&link).Error
	return
}

func (d *deviceRepoDriver) AddDeviceReview(r model.Review) (review model.Review, err error) {
	err = d.DB.Create(&r).Error
	return
}

func (d *deviceRepoDriver) AddDeviceTrait(deviceId uint, traitId uint) (device model.Device, err error) {
	trait := model.Trait{}
	traitErr := d.DB.Take(&trait, traitId).Error
	if traitErr != nil {
		err = traitErr
		return
	}
	err = d.DB.Model(&device).Association("DeviceTrait").Append(trait)
	return
}

func (d *deviceRepoDriver) DeleteDeviceLink(linkId uint) (err error) {
	err = d.DB.Delete(&model.Link{}, linkId).Error
	return
}

func (d *deviceRepoDriver) DeleteDeviceReview(r model.Review) (err error) {
	err = d.DB.Delete(&r).Error
	return
}

func (d *deviceRepoDriver) DeleteDeviceTrait(deviceId uint, traitId uint) (err error) {
	device := model.Device{}
	deviceErr := d.DB.Take(&device, deviceId).Error
	if deviceErr != nil {
		return deviceErr
	}

	trait := model.Trait{}
	traitErr := d.DB.Take(&trait, traitId).Error
	if traitErr != nil {
		return traitErr
	}

	err = d.DB.Model(&device).Association("DeviceBuy").Delete(trait)
	return
}

func (d *deviceRepoDriver) EditDeviceLink(b model.Link) (link model.Link, err error) {
	link = b
	err = d.DB.Save(&link).Error
	return
}

func (d *deviceRepoDriver) EditDeviceReview(r model.Review) (review model.Review, err error) {
	err = d.DB.Save(&r).Error
	return
}

func (d *deviceRepoDriver) EditDeviceScore(s model.Scores) (score model.Scores, err error) {
	score = s
	err = d.DB.Save(&score).Error
	return
}

func (d *deviceRepoDriver) GetDevice(deviceId uint) (device model.Device, err error) {
	err = d.DB.Preload(clause.Associations).Take(&device, deviceId).Error
	return
}

func (d *deviceRepoDriver) GetDeviceLink(deviceId uint, linkId uint) (link model.Link, err error) {
	err = d.DB.Where("id = ? AND device_id = ?", linkId, deviceId).Take(&link).Error
	return
}

func (d *deviceRepoDriver) GetDeviceLinks(deviceId uint) (links []*model.Link, err error) {
	device := model.Device{}
	err = d.DB.Preload("DeviceBuy").Find(&device, deviceId).Error
	links = device.Links
	return
}

func (d *deviceRepoDriver) GetDeviceReview(deviceId uint, reviewId uint) (review model.Review, err error) {
	err = d.DB.Where("id = ? AND device_id = ?", reviewId, deviceId).Take(&review).Error
	return
}

func (d *deviceRepoDriver) GetDeviceReviews(deviceId uint) (reviews []model.Review, err error) {
	err = d.DB.Where("device_id = ?", deviceId).Find(&reviews).Error
	return
}

func (d *deviceRepoDriver) GetDeviceScore(deviceId uint) (score model.Scores, err error) {
	err = d.DB.Where("device_id = ?", deviceId).Take(&score).Error
	return
}

func (d *deviceRepoDriver) GetDeviceTraits(deviceId uint) (traits []*model.Trait, err error) {
	device := model.Device{ID: deviceId}
	err = d.DB.Preload("Traits").Take(&device).Error
	traits = device.Traits
	return
}

func (d *deviceRepoDriver) GetDevices() (devices []model.Device, err error) {
	err = d.DB.Preload(clause.Associations).Omit("Category").Find(&devices).Error
	return
}

func NewDeviceRepoDriver(db *gorm.DB) domain.DeviceRepository {
	return &deviceRepoDriver{
		DB: db,
	}
}
