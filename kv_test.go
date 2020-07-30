package main_test

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"

	redimo "github.com/dbProjectRED/redimo"
	v1 "github.com/dbProjectRED/redimo/v1"
)

func TestRedimoServiceGetAndSet(t *testing.T) {
	service := redimo.RedimoService{}
	table := makeTable()
	value, _ := ptypes.MarshalAny(&wrappers.StringValue{Value: "v1"})
	_, err := service.Set(context.Background(), &v1.SetRequest{
		Table: &table,
		Key:   "k1",
		Value: value,
	})
	assert.NoError(t, err)

	resp, err := service.Get(context.Background(), &v1.GetRequest{
		Table: &table,
		Key:   "k1",
	})
	assert.NoError(t, err)

	sv := wrappers.StringValue{}
	assert.NoError(t, ptypes.UnmarshalAny(resp.Value, &sv))
	assert.Equal(t, "v1", sv.Value)
}

func makeTable() v1.Table {
	table := v1.Table{
		Table:              "testable",
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
		TableName:             aws.String("testable"),
		Tags:                  nil,
	}).Send(context.Background())

	return table
}
