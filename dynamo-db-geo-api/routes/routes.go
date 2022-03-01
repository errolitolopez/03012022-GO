package routes

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/errolitolopez/dynamo-db-geo-api/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(db *dynamodb.DynamoDB) *gin.Engine {
	routes := gin.Default()
	routes.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	routes.POST("/geoProvinces", controller.CreateGeoProvince)
	routes.GET("/geoProvinces", controller.GetAllGeoProvinces)
	routes.GET("/geoProvinces/:provinceCode", controller.GetOneByProvinceCode)
	routes.PATCH("/geoProvinces/:provinceCode", controller.UpdateGeoProvince)
	routes.DELETE("/geoProvinces/:provinceCode", controller.DeleteGeoProvince)

	routes.POST("/geoCities", controller.CreateGeoCities)
	routes.GET("/geoCities", controller.GetAllGeoCities)
	return routes
}
