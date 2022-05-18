package rest

import (
	"secondhand_glossary/internal/config"
	"secondhand_glossary/internal/domain"
)

type TokenController struct {
  conf config.Config
  s domain.TokenService
}
