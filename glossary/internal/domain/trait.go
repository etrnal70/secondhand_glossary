package domain

import "secondhand_glossary/internal/model"

type TraitRepository interface {
	GetTraits() (traits []model.Trait, err error)
	GetTraitDevices(traitId uint) (devices []model.Device, err error)
	AddTrait(t model.Trait) (trait model.Trait, err error)
	EditTrait(t model.Trait) (trait model.Trait, err error)
	DeleteTrait(traitId uint) (err error)
}

type TraitService interface {
	AddTrait(t model.Trait) (trait model.Trait, err error)
	EditTrait(t model.Trait) (trait model.Trait, err error) // Admin
	GetTraits() (traits []model.Trait, err error)
	GetTraitDevices(traitId uint) (devices []model.Device, err error)
	DeleteTrait(traitId uint) (err error) // Admin
}
