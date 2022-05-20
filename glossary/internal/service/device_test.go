package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDevice(t *testing.T) {
	expectedDevice := model.Device{Type: "Device Type"}

	mock := new(domain.MockDeviceRepository)
	mock.On("AddDevice", model.Device{Type: "Device Type"}).Return(expectedDevice, nil)

	deviceService := NewDeviceService(mock)

	device, err := deviceService.AddDevice(model.Device{Type: "Device Type"})

	assert.Nil(t, err)
	assert.Equal(t, expectedDevice.Type, device.Type)
}

func TestAddDeviceLink(t *testing.T) {
	expectedDevice := model.Link{Link: "example.com"}

	mock := new(domain.MockDeviceRepository)
	mock.On("AddDeviceLink", model.Link{Link: "example.com"}).Return(expectedDevice, nil)

	deviceService := NewDeviceService(mock)

	link, err := deviceService.AddDeviceLink(model.Link{Link: "example.com"})

	assert.Nil(t, err)
	assert.Equal(t, expectedDevice.Link, link.Link)
}

func TestAddDeviceReview(t *testing.T) {
	expectedReview := model.Review{UserID: 1, DeviceID: 1}

	mock := new(domain.MockDeviceRepository)
	mock.On("AddDeviceReview", model.Review{UserID: 1, DeviceID: 1}).Return(expectedReview, nil)

	deviceService := NewDeviceService(mock)

	newReview, err := deviceService.AddDeviceReview(model.Review{UserID: 1, DeviceID: 1})

	assert.NoError(t, err)
	assert.Equal(t, expectedReview.UserID, newReview.UserID)
	assert.Equal(t, expectedReview.DeviceID, newReview.DeviceID)
}

func TestAddDeviceTrait(t *testing.T) {
	expectedDevice := model.Device{
		ID: 1,
		Traits: []*model.Trait{
			{
				ID: 1,
			},
		},
	}

	mock := new(domain.MockDeviceRepository)
	mock.On("AddDeviceTrait", uint(1), uint(1)).Return(expectedDevice, nil)

	deviceService := NewDeviceService(mock)

	device, err := deviceService.AddDeviceTrait(1, 1)

	assert.NoError(t, err)
	assert.Equal(t, expectedDevice.Traits[0].ID, device.Traits[0].ID)
	assert.Equal(t, expectedDevice.ID, device.ID)
}
