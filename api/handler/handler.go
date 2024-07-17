package handler

import (
	"clent/generated/product"
	"clent/generated/users"

	"go.uber.org/zap"
)

type Handler struct {
	ProductClient product.ProductServiceClient
	UserClient users.AuthServiceClient
	Log *zap.Logger
}

func NewHandler(product product.ProductServiceClient, user users.AuthServiceClient,log *zap.Logger) *Handler {
	return &Handler{
		ProductClient: product,
		UserClient: user,
		Log:  log,
	}
}