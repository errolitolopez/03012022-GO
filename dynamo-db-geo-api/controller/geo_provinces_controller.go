package controller

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/errolitolopez/dynamo-db-geo-api/model"
	"github.com/gin-gonic/gin"
)

func GetAllGeoProvinces(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	params := &dynamodb.ScanInput{
		TableName: aws.String("GeoProvinces"),
	}
	resp, err := db.Scan(params)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var geoProvinces []model.GeoProvince
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &geoProvinces)
	c.JSON(http.StatusOK, gin.H{"data": geoProvinces})
}

func GetOneByProvinceCode(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	geoProvince := GetByProvinceCode(db, c.Param("provinceCode"))
	if geoProvince == (model.GeoProvince{}) {
		c.JSON(http.StatusOK, gin.H{"error": "No records found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": geoProvince})
}

func CreateGeoProvince(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	var geoProvince model.GeoProvince
	if err := c.ShouldBindJSON(&geoProvince); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if GetByProvinceCode(db, geoProvince.ProvinceCode) != (model.GeoProvince{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Already exists province code [%v]", geoProvince.ProvinceCode)})
		return
	}
	if provinceNameExists(db, geoProvince.ProvinceName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Already exists province name [%v]", geoProvince.ProvinceName)})
		return
	}
	geoProvinceMap, err := dynamodbattribute.MarshalMap(geoProvince)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("GeoProvinces"),
		Item:      geoProvinceMap,
	}
	resp, err := db.PutItem(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("resp: %v\n", resp)
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func UpdateGeoProvince(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	var geoProvince model.GeoProvince
	if err := c.ShouldBindJSON(&geoProvince); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	currentGeoProvince := GetByProvinceCode(db, c.Param("provinceCode"))
	if currentGeoProvince == (model.GeoProvince{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records found"})
		return
	}
	if currentGeoProvince.ProvinceName != geoProvince.ProvinceName && provinceNameExists(db, geoProvince.ProvinceName) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Already exists province name [%v]", geoProvince.ProvinceName)})
		return
	}
	params := &dynamodb.UpdateItemInput{
		TableName: aws.String("GeoProvinces"),
		Key: map[string]*dynamodb.AttributeValue{
			"ProvinceCode": {
				S: aws.String(currentGeoProvince.ProvinceCode),
			},
		},
		UpdateExpression: aws.String("SET ProvinceName = :provinceName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":provinceName": {
				S: aws.String(geoProvince.ProvinceName),
			},
		},
		ReturnValues: aws.String(dynamodb.ReturnValueAllNew),
	}
	resp, err := db.UpdateItem(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("resp: %v\n", resp)
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func DeleteGeoProvince(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	geoProvince := GetByProvinceCode(db, c.Param("provinceCode"))
	if geoProvince == (model.GeoProvince{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No records found"})
		return
	}
	params := &dynamodb.DeleteItemInput{
		TableName: aws.String("GeoProvinces"),
		Key: map[string]*dynamodb.AttributeValue{
			"ProvinceCode": {
				S: aws.String(geoProvince.ProvinceCode),
			},
		},
	}
	resp, err := db.DeleteItem(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("resp: %v\n", resp)
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func GetByProvinceCode(db *dynamodb.DynamoDB, provinceCode string) (geoProvince model.GeoProvince) {
	params := &dynamodb.ScanInput{
		TableName:        aws.String("GeoProvinces"),
		FilterExpression: aws.String(":provinceCode = ProvinceCode"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":provinceCode": {
				S: aws.String(provinceCode),
			},
		},
	}
	resp, _ := db.Scan(params)
	var geoProvinces []model.GeoProvince
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &geoProvinces)
	if len(geoProvinces) != 0 {
		geoProvince = geoProvinces[0]
	}
	return geoProvince
}

func provinceNameExists(db *dynamodb.DynamoDB, provinceName string) bool {
	params := &dynamodb.ScanInput{
		TableName:        aws.String("GeoProvinces"),
		FilterExpression: aws.String(":provinceName = ProvinceName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":provinceName": {
				S: aws.String(provinceName),
			},
		},
	}
	resp, _ := db.Scan(params)
	var geoProvinces []model.GeoProvince
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &geoProvinces)
	return len(geoProvinces) != 0
}
