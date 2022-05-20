package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddTrait(t *testing.T) {
	expectedTrait := model.Trait{Trait: "NewTrait"}

	mock := new(domain.MockTraitRepository)
	mock.On("AddTrait", model.Trait{Trait: "NewTrait"}).Return(expectedTrait, nil)

	traitService := NewTraitService(mock)

	_, err := traitService.AddTrait(expectedTrait)

	assert.Nil(t, err)
	assert.Equal(t, "NewTrait", expectedTrait.Trait)
}

func TestDeleteTrait(t *testing.T) {
	mock := new(domain.MockTraitRepository)
	mock.On("DeleteTrait", uint(1)).Return(nil)

	traitService := NewTraitService(mock)

	err := traitService.DeleteTrait(uint(1))

	assert.Nil(t, err)
}

func TestEditTrait(t *testing.T) {
	expectedTrait := model.Trait{Trait: "EditedTrait"}

	mock := new(domain.MockTraitRepository)
	mock.On("AddTrait", model.Trait{Trait: "EditedTrait"}).Return(expectedTrait, nil)

	traitService := NewTraitService(mock)

	_, err := traitService.AddTrait(expectedTrait)

	assert.Nil(t, err)
	assert.Equal(t, "EditedTrait", expectedTrait.Trait)
}

func TestGetTraitDevices(t *testing.T) {
	expectedDevices := []model.Device{
		{
			Traits: []*model.Trait{{ID: 1}},
		},
		{
			Traits: []*model.Trait{{ID: 1}},
		},
	}

	mock := new(domain.MockTraitRepository)
	mock.On("GetTraitDevices", uint(1)).Return(expectedDevices, nil)

	traitService := NewTraitService(mock)

	devices, err := traitService.GetTraitDevices(uint(1))

	assert.Nil(t, err)
	assert.Equal(t, expectedDevices[0].Traits[0].ID, devices[0].Traits[0].ID)
}

func TestGetTraits(t *testing.T) {
	expectedTraits := []model.Trait{{Trait: "Trait1"},{Trait: "Trait2"}}

	mock := new(domain.MockTraitRepository)
	mock.On("GetTraits").Return(expectedTraits, nil)

	traitService := NewTraitService(mock)

	traits, err := traitService.GetTraits()

	assert.Nil(t, err)
	assert.Equal(t, traits[0].Trait, expectedTraits[0].Trait)
}
