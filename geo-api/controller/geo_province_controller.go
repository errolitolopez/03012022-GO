package controller

import (
	"net/http"
	"time"

	"github.com/errolitolopez/geo-api/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// Get provinces
func FindGeoProvinces(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var geoProvinces []model.GeoProvince
	db.Find(&geoProvinces)
	c.JSON(http.StatusOK, gin.H{"data": geoProvinces})
}

// Get province by id
func FindGeoProvince(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var geoProvince model.GeoProvince
	if err := db.Where("id = ?", c.Param("id")).First(&geoProvince).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": geoProvince})
}

// Create province
func CreateGeoProvince(c *gin.Context) {
	var createGeoProvince model.CreateGeoProvince
	if err := c.ShouldBindJSON(&createGeoProvince); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	geoProvince := model.GeoProvince{
		Code:            createGeoProvince.Code,
		Name:            createGeoProvince.Name,
		Flag:            1,
		CreatedDate:     time.Now(),
		CreatedBy:       "USERMAIN",
		CreatedIp:       "0:0:0:0:0:0:0:1",
		LastUpdatedDate: time.Now(),
		LastUpdatedBy:   "USERMAIN",
		LastUpdatedIp:   "0:0:0:0:0:0:0:1",
	}
	db := c.MustGet("db").(*gorm.DB)
	db.Create(&geoProvince)
	c.JSON(http.StatusOK, gin.H{"data": geoProvince})
}

// Modify province
func ModifyCreateGeoProvince(c *gin.Context) {
	var modifyCreateGeoProvince model.ModifyCreateGeoProvince
	if err := c.ShouldBindJSON(&modifyCreateGeoProvince); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var geoProvince model.GeoProvince
	db := c.MustGet("db").(*gorm.DB)
	if err := db.Where("id = ?", modifyCreateGeoProvince.ID).First(&geoProvince).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	var updatedGeoProvince model.GeoProvince
	updatedGeoProvince.Code = modifyCreateGeoProvince.Code
	updatedGeoProvince.Name = modifyCreateGeoProvince.Name
	updatedGeoProvince.LastUpdatedBy = "SYSTEM"
	updatedGeoProvince.LastUpdatedIp = "192.168.0.1"
	updatedGeoProvince.LastUpdatedDate = time.Now()
	db.Model(&geoProvince).Updates(updatedGeoProvince)
	c.JSON(http.StatusOK, gin.H{"data": geoProvince})
}

// Delete province by id
func DeleteGeoProvince(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var geoProvince model.GeoProvince
	if err := db.Where("id = ?", c.Param("id")).First(&geoProvince).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	db.Delete(&geoProvince)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
