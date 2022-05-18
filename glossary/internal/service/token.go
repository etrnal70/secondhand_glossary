package service

import (
	"secondhand_glossary/internal/domain"
	"secondhand_glossary/internal/middleware/jwt"
)

type tokenService struct {
	Repo domain.TokenRepository
}

// CheckToken implements domain.TokenService
func (s *tokenService) CheckToken() bool {
	panic("unimplemented")
}

// RefreshAuthToken implements domain.TokenService
func (s *tokenService) RefreshAuthToken() (token jwt.TokenDetails, err error) {
	panic("unimplemented")
}

func NewTokenService(r domain.TokenRepository) domain.TokenService {
	return &tokenService{
		Repo: r,
	}
}
