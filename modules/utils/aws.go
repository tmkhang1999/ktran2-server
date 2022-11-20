package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"main.go/structs"
)

func CreateDynamoDBClient() *dynamodb.DynamoDB {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return dynamodb.New(sess)
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

func GetAll(svc *dynamodb.DynamoDB, tableName string) []structs.AwsItem {
	// get ScanOutput from dynamodb table
	collectionSO, _ := svc.Scan(&dynamodb.ScanInput{
		TableName: aws.String(tableName),
	})

	// Create the dynamic slice
	AwsItems := make([]structs.AwsItem, *collectionSO.Count)

	// Process collectionAV *dynamodb.ScanOutput
	for index, collectionsAV := range collectionSO.Items {
		AwsItem := structs.AwsItem{}

		err := dynamodbattribute.UnmarshalMap(collectionsAV, &AwsItem)
		if err != nil {
			log.Fatalln(err)
		}

		AwsItems[index] = AwsItem
	}

	return AwsItems
}
