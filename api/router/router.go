package router

import (
	_ "clent/api/docs"
	"clent/api/handler"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RouterApi @title API Service
// @version 1.0
// @description API service
// @host localhost:8085
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func RouterApi(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	products := router.Group("/api/v1/products")
	{
		products.POST("/product", h.CreateProductHandler)
		products.PUT("/:product_id", h.UpdateProductHandler)
		products.DELETE("/:product_id", h.DeleteProductHandler)
		products.GET("/get", h.GetProductHandler)
		products.GET("/:product_id", h.GetbyIdProductHandler)
		products.GET("/", h.GetallProducts)
		products.POST("/:product_id/rating", h.CreateRatingProductsHandler)
		products.GET("/:product_id/ratings", h.GetProductRatingsHandler)
	}

	order := router.Group("/api/v1/orders")
	{
		order.POST("/", h.CreateOrderHandler)
		order.DELETE("/:order_id", h.DeleteOrderHandler)
		order.PUT("/:order_id/update", h.UpdateOrderHandler)
		order.GET("/", h.GetallOrderHandler)
		order.GET("/:order_id", h.GetbyIdOrderHandler)
		order.POST("/:order_id/payment", h.CreatePaymentHandler)
		order.GET("/:order_id/payment-status", h.PaymentStatusHandler)
	}

	router.POST("/api/v1/product-categories", h.CreateCategoryHandler)
	router.GET("/api/v1/statistics", h.GetStatisticsHandler)
	router.GET("/api/v1/user-activity/:user_id", h.TrackUserActivityHandler)
	router.GET("/api/v1/recommendations", h.GetProductRecommendationsHandler)
	router.GET("/api/v1/artisan-rankings", h.GetanArtistRatingHandler)

	return router
}
