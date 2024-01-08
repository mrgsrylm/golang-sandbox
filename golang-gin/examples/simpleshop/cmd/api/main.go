package main

import (
	"github.com/mrgsrylm/simpleshop/pkg/config"
	"github.com/mrgsrylm/simpleshop/pkg/dbs"
	"github.com/mrgsrylm/simpleshop/pkg/logger"
	"github.com/mrgsrylm/simpleshop/pkg/redis"
	"github.com/mrgsrylm/simpleshop/pkg/validation"

	orderModel "github.com/mrgsrylm/simpleshop/internal/order/model"
	productModel "github.com/mrgsrylm/simpleshop/internal/product/model"
	grpcServer "github.com/mrgsrylm/simpleshop/internal/server/grpc"
	httpServer "github.com/mrgsrylm/simpleshop/internal/server/http"
	userModel "github.com/mrgsrylm/simpleshop/internal/user/model"
)

//	@title			simpleshop swagger api
//	@version		1.0
//	@description	Swagger API for simpleshop.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	Agus S. Mubarok
//	@contact.email	gusrylmubarok@gmail.com

//	@license.name	MIT
//	@license.url	https://github.com/mrgsrylm/simpleshop/blob/main/LICENSE

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name						Authorization

//	@BasePath	/api/v1
func main() {
	cfg := config.LoadConfig()
	logger.New(cfg.Environment)

	db, err := dbs.NewDatabase(cfg.DatabaseURI)
	if err != nil {
		logger.Fatal("Cannot connect to database", err)
	}

	err = db.AutoMigrate(&userModel.User{}, 
		&productModel.Product{}, 
		orderModel.Order{}, 
		orderModel.OrderLine{})
	if err != nil {
		logger.Fatal("Database migration fail", err)
	}

	validator := validation.New()

	cache := redis.New(redis.Config{
		Address:  cfg.RedisURI,
		Password: cfg.RedisPassword,
		Database: cfg.RedisDB,
	})

	go func() {
		httpSvr := httpServer.NewServer(validator, db, cache)
		if err = httpSvr.Run(); err != nil {
			logger.Fatal(err)
		}
	}()

	grpcSvr := grpcServer.NewServer(validator, db, cache)
	if err = grpcSvr.Run(); err != nil {
		logger.Fatal(err)
	}
}
