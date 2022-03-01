package main

import (
	"product-service/product"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/products", product.GetProducts)
	router.Run("localhost:9999")
}
