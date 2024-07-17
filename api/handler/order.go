package handler

import (
	pb "clent/generated/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateOrderHandler handles the creation of a new order.
// @Summary Create Order
// @Description Create a new order
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param Create body product.CreateOrderRequest true "Create Order"
// @Success 200 {object} product.CreateOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/ [post]
func (h *Handler) CreateOrderHandler(ctx *gin.Context) {
	var req pb.CreateOrderRequest

	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.ProductClient.CreateOrder(ctx, &req)
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

// DeleteOrderHandler handles the deletion of an order.
// @Summary Delete Order
// @Description Delete an existing order
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/{order_id} [delete]
func (h *Handler) DeleteOrderHandler(ctx *gin.Context) {
	id := ctx.Param("order_id")

	_, err := h.ProductClient.DeleteOrder(ctx, &pb.DeleteOrderRequest{OrderId: id})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, "ochirildi")
}

// UpdateOrderHandler handles the update of an order.
// @Summary Update Order
// @Description Update an existing order
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Param Update body product.UpdateOrderRequest true "Update Order"
// @Success 200 {object} product.UpdateOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/{order_id}/update [put]
func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	id := ctx.Param("order_id")
	var req pb.UpdateOrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	req.Id= id

	resp, err := h.ProductClient.UpdateOrder(ctx, &req)
	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}

// GetallOrderHandler retrieves a list of orders with optional filtering and pagination.
// @Summary Get All Orders
// @Description Retrieve a list of orders with optional filtering and pagination.
// @Tags Order
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param limit query int false "Number of items to return"
// @Param offset query int false "Offset for pagination"
// @Success 200 {object} product.GetallOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/ [get]
func (h *Handler) GetallOrderHandler(ctx *gin.Context) {
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

	resp, err := h.ProductClient.GetallOrder(ctx, &pb.GetallOrderRequest{
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

// GetbyIdOrderHandler handles the request to fetch an order by its ID.
// @Summary Get Order by ID
// @Description Get an order by its ID
// @Tags Order
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} product.GetByIdOrderResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/{order_id} [get]
func (h *Handler) GetbyIdOrderHandler(ctx *gin.Context) {
	id := ctx.Param("order_id")

	resp, err := h.ProductClient.GetByIdOrder(ctx, &pb.GetByIdOrderRerquest{OrderId: id})

	if err != nil {
		h.Log.Error(err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	h.Log.Info("ishladi")
	ctx.JSON(http.StatusOK, resp)
}
