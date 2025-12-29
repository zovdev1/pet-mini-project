package auth

import (
	"github.com/zovdev1/mini-app-project/internal/entites"
)

type TokenProvider interface {
	GenerateToken(user *entites.User) (string, error)
	ValidateToken(tokenString string) (*entites.User, error)
}
