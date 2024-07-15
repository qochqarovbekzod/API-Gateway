package handler

import (
	pb "clent/generated/product"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateOrderHandler(ctx *gin.Context) {
	var req pb.CreateOrderRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := h.ProductService.CreateOrder(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteOrderHandler(ctx *gin.Context) {
	id := ctx.Param("order_id")

	_, err := h.ProductService.DeleteOrder(ctx, &pb.DeleteOrderRequest{OrderId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, "ochirildi")

}

func (h *Handler) UpdateOrderHandler(ctx *gin.Context) {
	id := ctx.Param("order_id")
	var req pb.UpdateOrderRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	req.Id = id

	resp, err := h.ProductService.UpdateOrder(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) GetallOrderHandler(ctx *gin.Context) {
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

	resp, err := h.ProductService.GetallOrder(ctx, &pb.GetallOrderRequest{
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

func (h *Handler) GetbyIdOrderHandler(ctx *gin.Context) {
	id := ctx.Param("order_id")

	resp, err := h.ProductService.GetByIdOrder(ctx, &pb.GetByIdOrderRerquest{OrderId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}