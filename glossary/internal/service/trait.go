package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/model"
)

type traitService struct {
	Repo domain.TraitRepository
}

func (s *traitService) AddTrait(t model.Trait) (trait model.Trait, err error) {
	trait, err = s.Repo.AddTrait(t)
	return
}

func (s *traitService) DeleteTrait(traitId uint) (err error) {
	err = s.Repo.DeleteTrait(traitId)
	return
}

func (s *traitService) EditTrait(t model.Trait) (trait model.Trait, err error) {
	trait, err = s.Repo.EditTrait(t)
	return
}

func (s *traitService) GetTraitDevices(traitId uint) (devices []model.Device, err error) {
	devices, err = s.Repo.GetTraitDevices(traitId)
	return
}

func (s *traitService) GetTraits() (traits []model.Trait, err error) {
	traits, err = s.Repo.GetTraits()
	return
}

func NewTraitService(r domain.TraitRepository) domain.TraitService {
	return &traitService{
		Repo: r,
	}
}
