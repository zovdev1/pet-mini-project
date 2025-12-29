package v1

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/zovdev1/mini-app-project/internal/controller/rest/v1/request"
	"github.com/zovdev1/mini-app-project/internal/usecase/dto"
	"github.com/zovdev1/mini-app-project/pkg/logger"
	"go.uber.org/zap"
)

func (r *V1) Create(c *gin.Context) {
	id, err := GetUserId(c)
	startTime := time.Now()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var body request.Product

	if err := c.BindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	if err := r.v.Struct(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse request body"})
		return
	}

	pro := dto.ProductInput{
		Title:       body.Title,
		Description: body.Description,
		Price:       body.Price,
		Quantity:    body.Quantity,
		UserID:      id,
	}

	newProduct, err := r.p.Create(pro)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Product service problems"})
		return
	}

	c.JSON(http.StatusCreated, newProduct)

	logger.Info("Create",
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

func (r *V1) GetAllproduct(c *gin.Context) {

	startTime := time.Now()

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, _ := strconv.Atoi(limitStr)
	offser, _ := strconv.Atoi(offsetStr)

	product, err := r.p.GetAll(limit, offser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, product)

	logger.Info("GetAllproduct",
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

func (r *V1) GetListById(c *gin.Context) {
	id := c.Param("id")
	startTime := time.Now()

	product, err := r.p.GetById(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, product)

	logger.Info("GetListById",
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

func (r *V1) DeleteProduct(c *gin.Context) {
	userID, err := GetUserId(c)
	startTime := time.Now()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	productID := c.Param("id")

	err = r.p.DELETE(userID.String(), productID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		} else if strings.Contains(err.Error(), "invalid") {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok"})

	logger.Info("DeleteProduct",
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
