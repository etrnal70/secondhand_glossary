package repository

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"

	"gorm.io/gorm"
)

type traitRepoDriver struct {
	DB *gorm.DB
}

func (d *traitRepoDriver) AddTrait(t model.Trait) (trait model.Trait, err error) {
	trait = t
	err = d.DB.Save(&trait).Error
	return
}

func (d *traitRepoDriver) DeleteTrait(traitId uint) (err error) {
	var traitQuery *model.Trait
	queryErr := d.DB.Where("id = ?", traitId).Preload("Devices").First(&traitQuery).Error
	if queryErr != nil {
		err = queryErr
		return
	}

	deleteAssociationErr := d.DB.Model(&model.Trait{ID: traitId}).Association("Devices").Delete(&traitQuery.Devices)
	if deleteAssociationErr != nil {
		err = deleteAssociationErr
		return
	}

	err = d.DB.Delete(&model.Trait{}, traitId).Error
	return
}

func (d *traitRepoDriver) EditTrait(t model.Trait) (trait model.Trait, err error) {
	trait = t
	err = d.DB.Table("device_traits").Save(&trait).Error
	return
}

func (d *traitRepoDriver) GetTraitDevices(traitId uint) (devices []model.Device, err error) {
	err = d.DB.Model("device_traits").Where("device_trait_id = ?", traitId).Find(&devices).Error
	return
}

func (d *traitRepoDriver) GetTraits() (traits []model.Trait, err error) {
	err = d.DB.Find(&traits).Error
	return
}

func NewTraitRepoDriver(db *gorm.DB) domain.TraitRepository {
	return &traitRepoDriver{
		DB: db,
	}
}
