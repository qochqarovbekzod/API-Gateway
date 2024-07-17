package handler

import (
	pb "clent/generated/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRatingProductsHandler handles the creation of a new rating for a product.
// @Summary Create Rating
// @Description Create a new rating for a product
// @Tags Rating
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param product_id path string true "Product ID"
// @Param Create body product.CreateRatingProductsRequest true "Create Rating"
// @Success 200 {object} product.CreateRatingProductsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/{product_id}/rating [post]
func (h *Handler) CreateRatingProductsHandler(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	var req pb.CreateRatingProductsRequest

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	req.ProductId = productId

	resp, err := h.ProductClient.CreateRatingProducts(ctx, &req)

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetProductRatingsHandler handles the request to fetch ratings for a product by its ID.
// @Summary Get Ratings by Product ID
// @Description Get ratings for a product by its ID
// @Tags Rating
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param product_id path string true "Product ID"
// @Success 200 {object} product.GetProductRatingsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/{product_id}/ratings [get]
func (h *Handler) GetProductRatingsHandler(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	resp, err := h.ProductClient.GetProductRatings(ctx, &pb.GetProductRatingsRequest{ProductId: productId})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
