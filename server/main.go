package main

import (
	"clent/api/handler"
	"clent/api/router"
	"clent/config"
	"clent/generated/product"
	"clent/generated/users"
	 l"clent/logger"
	"fmt"
	"log"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var logger *zap.Logger

func initLog() {
	log, err := l.NewLogger()
	if err != nil {
		panic(err)
	}
	logger = log
}

func main() {
	initLog()
	cfg := config.Load()
	hand := NewConnect()

	router := router.RouterApi(hand)

	log.Fatal(router.Run(cfg.HTTP_PORT))
}

func NewConnect() *handler.Handler {
	usersConn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}

	usersService := users.NewAuthServiceClient(usersConn)

	productConn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Println("error")
		log.Println("ERROR: ", err.Error())
		return nil
	}

	productService := product.NewProductServiceClient(productConn)

	return &handler.Handler{
		UserClient:    usersService,
		ProductClient: productService,
		Log:           logger,
	}
}
