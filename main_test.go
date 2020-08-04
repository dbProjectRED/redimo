package main_test

import (
	"context"
	"crypto/rand"
	"math/big"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	v1 "github.com/dbProjectRED/redimo/v1"
)

func randSeq(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)

	for i := range b {
		n2, _ := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		b[i] = letters[n2.Int64()]
	}

	return string(b)
}

func makeTable() v1.Table {
	table := v1.Table{
		Table:              randSeq(10),
		AwsAccessKeyId:     "abc",
		AwsSecretAccessKey: "def",
		AwsSessionToken:    "ghi",
		AwsRegion:          "ap-south-1",
		EndpointOverride:   "http://localhost:8000",
	}
	_, _ = table.Client().CreateTableRequest(&dynamodb.CreateTableInput{
		AttributeDefinitions: []dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("pk"),
				AttributeType: dynamodb.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("sk"),
				AttributeType: dynamodb.ScalarAttributeTypeS,
			},
		},
		BillingMode:            dynamodb.BillingModePayPerRequest,
		GlobalSecondaryIndexes: nil,
		KeySchema: []dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("pk"),
				KeyType:       dynamodb.KeyTypeHash,
			},
			{
				AttributeName: aws.String("sk"),
				KeyType:       dynamodb.KeyTypeRange,
			},
		},
		LocalSecondaryIndexes: nil,
		ProvisionedThroughput: nil,
		SSESpecification:      nil,
		StreamSpecification:   nil,
		TableName:             aws.String(table.GetTable()),
		Tags:                  nil,
	}).Send(context.Background())

	return table
}
