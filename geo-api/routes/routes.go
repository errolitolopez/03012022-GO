package routes

import (
	"github.com/errolitolopez/geo-api/controller"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func SetupRoutes(db *gorm.DB) *gin.Engine {
	routes := gin.Default()
	routes.Use(func(c *gin.Context) {
		c.Set("db", db)
	})
	routes.POST("/geoProvinces", controller.CreateGeoProvince)
	routes.GET("/geoProvinces", controller.FindGeoProvinces)
	routes.GET("/geoProvinces/:id", controller.FindGeoProvince)
	routes.PATCH("/geoProvinces", controller.ModifyCreateGeoProvince)
	routes.DELETE("/geoProvinces/:id", controller.DeleteGeoProvince)
	return routes
}
