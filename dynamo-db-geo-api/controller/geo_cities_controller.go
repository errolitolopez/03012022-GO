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

func GetAllGeoCities(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	params := &dynamodb.ScanInput{
		TableName: aws.String("GeoCities"),
	}
	resp, err := db.Scan(params)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	var geoCities []model.GeoCity
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &geoCities)
	c.JSON(http.StatusOK, gin.H{"data": geoCities})
}
func CreateGeoCities(c *gin.Context) {
	db := c.MustGet("db").(*dynamodb.DynamoDB)
	var geoCity model.GeoCity
	if err := c.ShouldBindJSON(&geoCity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	geoProvince := GetByProvinceCode(db, geoCity.ProvinceCode)
	if geoProvince == (model.GeoProvince{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("No records found for province code [%v]", geoCity.ProvinceCode)})
		return
	}
	if GetByCityCode(db, geoCity.CityCode) != (model.GeoCity{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Already exists city code [%v]", geoCity.CityCode)})
		return
	}
	if cityNameExistsWithinProvince(db, geoCity) {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Already exists city name [%v] in province [%v]", geoCity.CityName, geoProvince.ProvinceName)})
		return
	}
	geoCityMap, err := dynamodbattribute.MarshalMap(geoCity)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	params := &dynamodb.PutItemInput{
		TableName: aws.String("GeoCities"),
		Item:      geoCityMap,
	}
	resp, err := db.PutItem(params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Printf("resp: %v\n", resp)
	c.JSON(http.StatusOK, gin.H{"message": "Success"})
}

func GetByCityCode(db *dynamodb.DynamoDB, cityCode string) (geoCity model.GeoCity) {
	params := &dynamodb.ScanInput{
		TableName:        aws.String("GeoCities"),
		FilterExpression: aws.String(":cityCode = CityCode"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":cityCode": {
				S: aws.String(cityCode),
			},
		},
	}
	resp, _ := db.Scan(params)
	var geoCities []model.GeoCity
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &geoCities)
	if len(geoCities) != 0 {
		geoCity = geoCities[0]
	}
	return geoCity
}

func cityNameExistsWithinProvince(db *dynamodb.DynamoDB, geoCity model.GeoCity) bool {
	params := &dynamodb.ScanInput{
		TableName:        aws.String("GeoCities"),
		FilterExpression: aws.String(":provinceCode = ProvinceCode AND :cityName = CityName"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":provinceCode": {
				S: aws.String(geoCity.ProvinceCode),
			},
			":cityName": {
				S: aws.String(geoCity.CityName),
			},
		},
	}
	resp, _ := db.Scan(params)
	var geoProvinces []model.GeoProvince
	_ = dynamodbattribute.UnmarshalListOfMaps(resp.Items, &geoProvinces)
	return len(geoProvinces) != 0
}
