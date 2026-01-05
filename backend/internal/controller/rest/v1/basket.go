package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zovdev1/mini-app-project/internal/controller/rest/v1/request"
)

func (r *V1) AddItem(c *gin.Context) {
	idUser, err := GetUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var input request.Basket

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	go func() {
		err = r.b.BasketAdd(idUser, input.Product_id)
	}()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "basket service problems",
			"info":  err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (r *V1) GetBasket(c *gin.Context) {
	idUser, err := GetUserId(c)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	resource, err := r.b.BasketGet(idUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch Basket_Item"})
		return
	}

	c.JSON(http.StatusOK, resource)
}
