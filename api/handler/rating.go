package handler

import (
	pb "clent/generated/product"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateRatingProductsHandler(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	var req pb.CreateRatingProductsRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	req.ProductId = productId

	resp, err := h.ProductService.CreateRatingProducts(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetProductRatingsHandler(ctx *gin.Context) {
	productId := ctx.Param("product_id")

	resp, err := h.ProductService.GetProductRatings(ctx, &pb.GetProductRatingsRequest{ProductId: productId})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK,resp)
}

