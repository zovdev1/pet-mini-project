package rest

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/zovdev1/mini-app-project/internal/controller/rest/v1"
	"github.com/zovdev1/mini-app-project/internal/usecase"
	auth "github.com/zovdev1/mini-app-project/pkg/jwt"
)

func NewRouter(app *gin.Engine, u usecase.User, p usecase.Product, b usecase.Basket, j auth.TokenProvider) {

	apiV1Group := app.Group("/api/v1")
	{
		v1.NewUser(apiV1Group, u, p, b, j)
	}
}
