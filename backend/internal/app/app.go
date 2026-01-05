package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"
	"github.com/zovdev1/mini-app-project/internal/controller/rest"
	persistent "github.com/zovdev1/mini-app-project/internal/repo/postgresql"
	redisdb "github.com/zovdev1/mini-app-project/internal/repo/redisDB"
	"github.com/zovdev1/mini-app-project/internal/usecase/basket"
	"github.com/zovdev1/mini-app-project/internal/usecase/product"
	"github.com/zovdev1/mini-app-project/internal/usecase/user"
	auth "github.com/zovdev1/mini-app-project/pkg/jwt"
	"github.com/zovdev1/mini-app-project/pkg/logger"
	"go.uber.org/zap"
)

var (
	secretKey = os.Getenv("JWT_SECRET_KEY")
	ctx       = context.Background()
)

const Time = 24 * time.Hour

func Run() {

	err := gotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	defer logger.Sync()

	config, err := persistent.NewPostgresDB(persistent.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	client, err := redisdb.NewConfig(ctx, redisdb.Config{
		Addr:        os.Getenv("REDIS_PORT"),
		Password:    "",
		User:        "user",
		DB:          0,
		MaxRetries:  3,
		DialTimeout: 5 * time.Second,
		Timeout:     3 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	// infrastructure
	redisProduct := redisdb.NewRedisUseRepo(client)
	repoUser := persistent.NewUseRepo(config)
	repoProduct := persistent.NewProductUseRepo(config)
	repoBasket := persistent.NewBasketUseRepo(config)
	repoBasketItem := persistent.NewBasketTimeUseRepo(config)
	authUser := auth.NewTokenJwt(secretKey, Time)

	// UseCase
	UserUseCase := user.New(repoUser, authUser)
	ProductUseCase := product.New(repoProduct, redisProduct)
	BasketUseCase := basket.New(repoBasket, repoBasketItem)

	// rest - api
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
		},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	rest.NewRouter(router, UserUseCase, ProductUseCase, BasketUseCase, authUser)

	logger.Info("Listening and serving HTTP", zap.String("port", os.Getenv("PORT")))
	router.Run(":" + os.Getenv("PORT"))
}
