package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (r *V1) GinserIdentity(c *gin.Context) {
	hendler := c.GetHeader(authorizationHeader)

	if hendler == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "empty auth header"})
		return
	}

	headerParts := strings.Split(hendler, " ")

	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid auth header"})
		return
	}

	if len(headerParts[1]) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "token is empty"})
		return
	}

	userId, err := r.j.ValidateToken(headerParts[1])

	if err != nil {
		c.JSON(http.StatusUnauthorized, err.Error())
		return
	}

	c.Set(userCtx, userId.ID)
}

func GetUserId(c *gin.Context) (uuid.UUID, error) {
	id, ok := c.Get(userCtx)

	if !ok {
		return uuid.Nil, errors.New("user id not found")
	}

	iduu, ok := id.(uuid.UUID)

	if !ok {
		return uuid.Nil, errors.New("user id is of invalid type")
	}

	return iduu, nil
}
