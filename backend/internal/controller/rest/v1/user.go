package v1

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zovdev1/mini-app-project/internal/controller/rest/v1/request"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
	"github.com/zovdev1/mini-app-project/pkg/logger"
	"go.uber.org/zap"
)

func (r *V1) SignUp(c *gin.Context) {

	startTime := time.Now()

	var body request.User

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	if err := r.v.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	createReq := dto.RegisterInput{
		Email:    body.Email,
		Password: body.Password,
	}

	newUserUUID, err := r.u.User(createReq)

	if err != nil {

		if strings.Contains(err.Error(), "already registered") {
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}

		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User service problems: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newUserUUID)

	logger.Info("SignUp",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.String("protocol", c.Request.Proto),
		zap.String("host", c.Request.Host),
		zap.Bool("tls", c.Request.TLS != nil),
		zap.String("user_agent", c.GetHeader("User-Agent")),
		zap.String("referer", c.GetHeader("Referer")),
		zap.String("accept_language", c.GetHeader("Accept-Language")),
		zap.String("ip", c.ClientIP()),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("duration", time.Since(startTime)),
	)
}

func (r *V1) LogIn(c *gin.Context) {

	startTime := time.Now()

	var body dto.SignInInput

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	if err := r.v.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	token, err := r.u.LogInUser(body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Token service problems"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

	logger.Info("LogIn",
		zap.String("method", c.Request.Method),
		zap.String("path", c.Request.URL.Path),
		zap.String("protocol", c.Request.Proto),
		zap.String("host", c.Request.Host),
		zap.Bool("tls", c.Request.TLS != nil),
		zap.String("user_agent", c.GetHeader("User-Agent")),
		zap.String("referer", c.GetHeader("Referer")),
		zap.String("accept_language", c.GetHeader("Accept-Language")),
		zap.String("ip", c.ClientIP()),
		zap.Int("status", c.Writer.Status()),
		zap.Duration("duration", time.Since(startTime)),
	)
}
