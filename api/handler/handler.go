package handler

import "clent/generated/product"

type Handler struct {
	ProductService product.ProductServiceClient
}

func NewHandler(product product.ProductServiceClient) *Handler {
	return &Handler{
		ProductService: product,
	}
}