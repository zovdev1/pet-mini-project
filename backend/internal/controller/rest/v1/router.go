package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/zovdev1/mini-app-project/internal/usecase"
	auth "github.com/zovdev1/mini-app-project/pkg/jwt"
)

func NewUser(apiV1Group *gin.RouterGroup, u usecase.User, p usecase.Product, b usecase.Basket, j auth.TokenProvider) {
	router := &V1{
		u: u,
		p: p,
		b: b,
		j: j,
		v: validator.New(validator.WithRequiredStructEnabled()),
	}

	UserGroup := apiV1Group.Group("/user")
	{
		UserGroup.POST("/create", router.SignUp)
		UserGroup.POST("/logIn", router.LogIn)
	}

	ProductGroup := apiV1Group.Group("/product")
	{
		ProductGroup.POST("/", router.GinserIdentity, router.Create)
		ProductGroup.GET("/", router.GetAllproduct)
		ProductGroup.GET("/:id", router.GetListById)
		ProductGroup.DELETE("/:id", router.GinserIdentity, router.DeleteProduct)
	}

	BasketGroup := apiV1Group.Group("/basket")
	{
		BasketGroup.POST("/items", router.GinserIdentity, router.AddItem) //  добавить товар и создать корзину
		BasketGroup.GET("/all", router.GinserIdentity, router.GetBasket)  // получить корзину
		// BasketGroup.DELETE("/item/:id")  // удалить по id
		// BasketGroup.PUT("/basket/items") // изменить количество
	}
}
