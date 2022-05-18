package domain

import "secondhand_glossary/internal/middleware/jwt"

type TokenRepository interface {
	SetToken(userId uint, token jwt.TokenDetails) (err error)
	SetRefreshToken()
	DeleteToken(id string) (int64, error)
	GetToken(td jwt.TokenDetails) (uint64, error)
}

type TokenService interface {
	RefreshAuthToken() (token jwt.TokenDetails, err error)
	CheckToken() bool
}
