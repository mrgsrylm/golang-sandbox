package http

import (
	"github.com/gin-gonic/gin"
	"github.com/mrgsrylm/simpleshop/internal/product/repository"
	"github.com/mrgsrylm/simpleshop/internal/product/service"
	"github.com/mrgsrylm/simpleshop/pkg/dbs"
	"github.com/mrgsrylm/simpleshop/pkg/middleware"
	"github.com/mrgsrylm/simpleshop/pkg/redis"
	"github.com/mrgsrylm/simpleshop/pkg/validation"
)

func Routes(r *gin.RouterGroup, db dbs.IDatabase, validator validation.Validation, cache redis.IRedis) {
	productRepo := repository.NewProductRepository(db)
	productSvc := service.NewProductService(validator, productRepo)
	productHandler := NewProductHandler(cache, productSvc)

	authMiddleware := middleware.JWTAuth()

	productRoute := r.Group("/products")
	{
		productRoute.GET("", productHandler.ListProducts)
		productRoute.POST("", authMiddleware, productHandler.CreateProduct)
		productRoute.PUT("/:id", authMiddleware, productHandler.UpdateProduct)
		productRoute.GET("/:id", productHandler.GetProductByID)
	}
}
