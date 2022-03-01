package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	session, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-east-2"),
		Credentials: credentials.NewSharedCredentials("", "test-account"),
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	db := dynamodb.New(session)

	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Movies"),
		KeySchema: []*dynamodb.KeySchemaElement{
			{AttributeName: aws.String("year"), KeyType: aws.String("HASH")},
			{AttributeName: aws.String("title"), KeyType: aws.String("RANGE")},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{AttributeName: aws.String("year"), AttributeType: aws.String("N")},
			{AttributeName: aws.String("title"), AttributeType: aws.String("S")},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(100),
		},
	}

	resp, err := db.CreateTable(params)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(resp)
}
