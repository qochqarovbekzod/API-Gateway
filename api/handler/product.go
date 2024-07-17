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

// CreateProductHandler handles the creation of a new menu item.
// @Summary Create Product 
// @Description Create a new Product item
// @Tags Product
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body product.CreateProductRequest true "Create Product"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/product [post]
func (h *Handler) CreateProductHandler(ctx *gin.Context) {
	var req pb.CreateProductRequest

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if Parse(req.CategoryId) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "CategoryId hato",
		})
	}

	resp, err := h.ProductClient.CreateProduct(ctx, &req)

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// UpdateProductHandler handles the update of a menu item.
// @Summary Update Product
// @Description Update an existing product item
// @Tags Product
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Update body product.UpdateProductRequest true "Update Menu"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/{product_id} [put]
func (h *Handler) UpdateProductHandler(ctx *gin.Context) {
	var req pb.UpdateProductRequest
	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if Parse(req.Id) {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "CategoryId hato",
		})
	}

	resp, err := h.ProductClient.UpdateProduct(ctx, &req)
	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)

}

// DeleteProductHandler handles the deletion of a menu item.
// @Summary Delete Product
// @Description Delete an existing product item
// @Tags Product
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/{product_id} [delete]
func (h *Handler) DeleteProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	_, err := h.ProductClient.DeleteProduct(ctx, &pb.DeleteProductRequest{Id: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, "ochirildi")

}
// GetProductHandler retrieves a list of orders with optional filtering and pagination.
// @Summary Get Product
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "limit of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} product.GetProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/get [get]
func (h *Handler) GetProductHandler(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offset := ctx.Query("offset")
	var limit2, offset2 int
	var err error

	if limit != "" {
		limit2, err = strconv.Atoi(limit)
		if err != nil {
			h.Log.Error(err.Error())
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
			h.Log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	resp, err := h.ProductClient.GetProduct(ctx, &pb.GetProductRequest{
		Limit:  int32(limit2),
		Offset: int32(offset2),
	})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetbyIdProductHandler handles the request to fetch a menu item by its ID.
// @Summary Get Product by ID
// @Description Get a Product item by its ID
// @Tags Product
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID"
// @Success 200 {object} product.GetbyIdProductResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/{product_id} [get]
func (h *Handler) GetbyIdProductHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	resp, err := h.ProductClient.GetbyIdProduct(ctx, &pb.GetbyIdProductRequest{Id: id})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetallProducts retrieves a list of menu items with optional filtering and pagination.
// @Summary Get All Product
// @Description Retrieve a list of Product items with optional filtering and pagination.
// @Tags Product
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param name query string false "Filter by Product item name"
// @Param description query string false "Filter by Product item description"
// @Param restaurant_id query string false "Filter by restaurant ID"
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Param price query string false "Filter by Product item price"
// @Success 200 {object} product.GetallProductsResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/products/ [get]
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
			h.Log.Error(err.Error())
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
			h.Log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	if max_price != "" {
		max_price2, err = strconv.ParseFloat(max_price, 64)
		if err != nil {
			h.Log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	if min_price != "" {
		min_price2, err = strconv.ParseFloat(min_price, 64)
		if err != nil {
			h.Log.Error(err.Error())
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
	}

	resp, err := h.ProductClient.GetallProducts(ctx, &pb.GetallProductsRequest{
		Category: category,
		MinPrice: float32(min_price2),
		MaxPrice: float32(max_price2),
		Limit:    int32(limit2),
		Offset:   int32(offset2),
	})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
