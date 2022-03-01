package product

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID    int     `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func GetProducts(c *gin.Context) {
	products := []Product{
		{
			1,
			"Book",
			1.99,
		},
		{
			2,
			"Carpet",
			59.99,
		},
	}
	c.IndentedJSON(http.StatusOK, products)
}

func NewProduct(id int, title string, price float64) Product {
	return Product{id, title, price}
}
