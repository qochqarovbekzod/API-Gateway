package handler

import (
	pb "clent/generated/product"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePaymentHandler(ctx *gin.Context) {
	orderId := ctx.Param("order_id")

	var req pb.CreatePaymentRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if req.PaymentMethod != "credit_card" && req.PaymentMethod != "cash" && req.PaymentMethod != "paymend" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "tolov turi notogri",
		})
	}

	if len(req.CardNumber) == 16 && req.PaymentMethod == "credit_card" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "card raqam hato",
		})
	}

	if len(req.ExpiryDate) == 15 {
		tel := strings.Split(req.ExpiryDate, "/")
		for _, v := range tel {
			_, err := strconv.Atoi(v)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"error": "ExpiryDate notogri",
				})
				return
			}
		}
	}

	if len(req.Svv) == 3 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "svv hato",
		})
		return
	}

	req.OrderId = orderId

	resp, err := h.ProductService.CreatePayment(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) PaymentStatusHandler(ctx *gin.Context) {
	orderId := ctx.Param("order_id")

	resp, err := h.ProductService.PaymentStatus(ctx, &pb.PaymentStatusRequest{OrderId: orderId})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
