package handler

import (
	pb "clent/generated/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Parse(id string) bool {
	_, err := uuid.Parse(id)
	return !(err == nil)
}

func (h *Handler) CreateProductHandler(ctx *gin.Context) {
	var req pb.CreateProductRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if Parse(req.CategoryId) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "CategoryId hato",
		})
	}

	resp, err := h.ProductService.CreateProduct(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateProductHandler(ctx *gin.Context) {
	var req pb.UpdateProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if Parse(req.Id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "CategoryId hato",
		})
	}

	resp, err := h.ProductService.UpdateProduct(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) DeleteProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := h.ProductService.DeleteProduct(ctx, &pb.DeleteProductRequest{Id: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, "ochirildi")

}

func (h *Handler) GetProductHandler(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	var limit2, offset2 int
	var err error

	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		limit2 = 5
	}

	if offset != "" {
		offset2, err = strconv.Atoi(offset)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	resp, err := h.ProductService.GetProduct(ctx, &pb.GetProductRequest{
		Limit:  int32(limit2),
		Offset: int32(offset2),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetbyIdProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.ProductService.GetbyIdProduct(ctx, &pb.GetbyIdProductRequest{Id: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetallProducts(ctx *gin.Context) {
	category := ctx.Query("category")
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	min_price := ctx.Query("min_price")
	max_price := ctx.Query("max_price")
	var limit2, offset2 int
	var min_price2, max_price2 float64
	var err error

	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	} else {
		limit2 = 5
	}

	if offset != "" {
		offset2, err = strconv.Atoi(offset)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	if max_price != "" {
		max_price2, err = strconv.ParseFloat(max_price, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	if min_price != "" {
		min_price2, err = strconv.ParseFloat(min_price, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	resp, err := h.ProductService.GetallProducts(ctx, &pb.GetallProductsRequest{
		Category: category,
		MinPrice: float32(min_price2),
		MaxPrice: float32(max_price2),
		Limit:    int32(limit2),
		Offset:   int32(offset2),
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}
