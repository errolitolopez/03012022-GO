package db

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func SetupDB() *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("", "test-account"),
	})
	if err != nil {
		fmt.Printf("err.Error(): %v\n", err.Error())
		panic(err.Error())
	}
	db := dynamodb.New(sess)
	// createTable(db)
	return db
}

// func createTable(db *dynamodb.DynamoDB) {
// 	tableName := "GeoProvinces"
// 	params := &dynamodb.CreateTableInput{
// 		TableName: aws.String(tableName),
// 		KeySchema: []*dynamodb.KeySchemaElement{
// 			{AttributeName: aws.String("ProvinceCode"), KeyType: aws.String("HASH")},
// 		},
// 		AttributeDefinitions: []*dynamodb.AttributeDefinition{
// 			{AttributeName: aws.String("ProvinceCode"), AttributeType: aws.String("S")},
// 		},
// 		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
// 			ReadCapacityUnits:  aws.Int64(10),
// 			WriteCapacityUnits: aws.Int64(100),
// 		},
// 	}
// 	_, err := db.CreateTable(params)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		return
// 	}
// }

// func deleteTable(db *dynamodb.DynamoDB, tableName string) {
// 	params := &dynamodb.DeleteTableInput{
// 		TableName: aws.String(tableName),
// 	}
// 	_, err := db.DeleteTable(params)
// 	if err != nil {
// 		fmt.Printf("err: %v\n", err)
// 		return
// 	}
// 	fmt.Printf("Table deleted: %v\n", tableName)
// }
