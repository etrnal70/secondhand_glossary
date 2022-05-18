package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
)

type deviceService struct {
	Repo domain.DeviceRepository
}

func (s *deviceService) AddDevice(d model.Device) (device model.Device, err error) {
	device, err = s.Repo.AddDevice(d)
	return
}

func (s *deviceService) AddDeviceLink(b model.Link) (link model.Link, err error) {
	link, err = s.Repo.AddDeviceLink(b)
	return
}

func (s *deviceService) AddDeviceReview(r model.Review) (review model.Review, err error) {
	review, err = s.Repo.AddDeviceReview(r)
	return
}

func (s *deviceService) AddDeviceTrait(deviceId uint, traitId uint) (device model.Device, err error) {
	device, err = s.Repo.AddDeviceTrait(deviceId, traitId)
	return
}

func (s *deviceService) DeleteDevice(deviceId uint) (err error) {
	err = s.Repo.DeleteDevice(deviceId)
	return
}

func (s *deviceService) DeleteDeviceLink(linkId uint) (err error) {
	err = s.Repo.DeleteDeviceLink(linkId)
	return
}

func (s *deviceService) DeleteDeviceReview(r model.Review) (err error) {
	err = s.Repo.DeleteDeviceReview(r)
	return
}

func (s *deviceService) DeleteDeviceTrait(deviceId uint, traitId uint) (err error) {
	err = s.Repo.DeleteDeviceTrait(deviceId, traitId)
	return
}

func (s *deviceService) EditDevice(d model.Device) (device model.Device, err error) {
	device, err = s.Repo.EditDevice(d)
	return
}

func (s *deviceService) EditDeviceLink(b model.Link) (link model.Link, err error) {
	link, err = s.Repo.EditDeviceLink(b)
	return
}

func (s *deviceService) EditDeviceReview(r model.Review) (review model.Review, err error) {
	review, err = s.Repo.EditDeviceReview(r)
	return
}

func (s *deviceService) GetDevice(deviceId uint) (device model.Device, err error) {
	device, err = s.Repo.GetDevice(deviceId)
	return
}

func (s *deviceService) GetDeviceLink(deviceId uint, linkId uint) (link model.Link, err error) {
	link, err = s.Repo.GetDeviceLink(deviceId, linkId)
	return
}

func (s *deviceService) GetDeviceLinks(deviceId uint) (links []*model.Link, err error) {
	links, err = s.Repo.GetDeviceLinks(deviceId)
	return
}

func (s *deviceService) GetDeviceReview(deviceId uint, reviewId uint) (review model.Review, err error) {
	review, err = s.Repo.GetDeviceReview(deviceId, reviewId)
	return
}

func (s *deviceService) GetDeviceReviews(deviceId uint) (reviews []model.Review, err error) {
	reviews, err = s.Repo.GetDeviceReviews(deviceId)
	return
}

func (s *deviceService) GetDeviceScore(deviceId uint) (score model.Scores, err error) {
	score, err = s.Repo.GetDeviceScore(deviceId)
	return
}

func (s *deviceService) GetDeviceTraits(deviceId uint) (traits []*model.Trait, err error) {
	traits, err = s.Repo.GetDeviceTraits(deviceId)
	return
}

func (s *deviceService) GetDevices() (devices []model.Device, err error) {
	devices, err = s.Repo.GetDevices()
	return
}

func (s *deviceService) EditDeviceScore(scr model.Scores) (score model.Scores, err error) {
	score, err = s.Repo.EditDeviceScore(scr)
	return
}

func NewDeviceService(r domain.DeviceRepository) domain.DeviceService {
	return &deviceService{
		Repo: r,
	}
}
