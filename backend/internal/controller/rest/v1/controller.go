package v1

import (
	"github.com/go-playground/validator/v10"
	"github.com/zovdev1/mini-app-project/internal/usecase"
	auth "github.com/zovdev1/mini-app-project/pkg/jwt"
)

type V1 struct {
	u usecase.User
	p usecase.Product
	j auth.TokenProvider
	v *validator.Validate
}
