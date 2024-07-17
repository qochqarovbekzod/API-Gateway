package handler

import (
    pb "clent/generated/product"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
)

// CreatePaymentHandler handles the creation of a new payment.
// @Summary Create Payment
// @Description Create a new payment
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param order_id path string true "Order ID"
// @Param Create body product.CreatePaymentRequest true "Create Payment"
// @Success 200 {object} product.CreatePaymentResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/{order_id}/payment [post]
func (h *Handler) CreatePaymentHandler(ctx *gin.Context) {
    orderId := ctx.Param("order_id")

    var req pb.CreatePaymentRequest

    if err := ctx.ShouldBind(&req); err != nil {
        h.Log.Error(err.Error())
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        return
    }

    if req.PaymentMethod != "credit_card" && req.PaymentMethod != "cash" && req.PaymentMethod != "paymend" {
       h.Log.Error("tolov tiri hato")
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "tolov turi notogri",
        })
        return
    }

    if len(req.CardNumber) != 16 && req.PaymentMethod == "credit_card" {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "card raqam hato",
        })
        return
    }

    if len(req.ExpiryDate) != 5 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "ExpiryDate notogri",
        })
        return
    }

    expiryParts := strings.Split(req.ExpiryDate, "/")
    if len(expiryParts) != 2 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "ExpiryDate notogri formatda",
        })
        return
    }

    for _, part := range expiryParts {
        if _, err := strconv.Atoi(part); err != nil {
            h.Log.Error(err.Error())
            ctx.JSON(http.StatusBadRequest, gin.H{
                "error": "ExpiryDate notogri",
            })
            return
        }
    }

    if len(req.Svv) != 3 {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "error": "svv hato",
        })
        return
    }

    req.OrderId = orderId

    resp, err := h.ProductClient.CreatePayment(ctx, &req)
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

// PaymentStatusHandler handles checking the status of a payment.
// @Summary Payment Status
// @Description Check the status of a payment
// @Tags Payment
// @Accept json
// @Security BearerAuth
// @Produce json
// @Param order_id path string true "Order ID"
// @Success 200 {object} product.PaymentStatusResponse
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /api/v1/orders/{order_id}/payment-status [get]
func (h *Handler) PaymentStatusHandler(ctx *gin.Context) {
    orderId := ctx.Param("order_id")

    resp, err := h.ProductClient.PaymentStatus(ctx, &pb.PaymentStatusRequest{OrderId: orderId})
    if err != nil {
        h.Log.Error(err.Error())
        ctx.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    ctx.JSON(http.StatusOK, resp)
}
