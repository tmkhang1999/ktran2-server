package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"main.go/structs"
)

func CreateDynamoDBClient(region string) *dynamodb.DynamoDB {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)
	if err != nil {
		log.Fatalf("Got error initializing AWS: %s", err)
	}
	svc := dynamodb.New(sess)

	return svc
}

func GetStatus(svc *dynamodb.DynamoDB, tableName string) structs.HttpStatusMessage {
	// get ScanOutput from dynamodb table
	collectionSO, _ := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	// init HttpStatusMessage
	httpStatusMessage := structs.HttpStatusMessage{
		Table:       tableName,
		RecordCount: collectionSO.Count,
	}

	return httpStatusMessage
}
