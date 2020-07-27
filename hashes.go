package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/golang/protobuf/ptypes/any"

	v1 "github.com/dbProjectRED/redimo/v1"
)

func (r RedimoService) HGET(ctx context.Context, request *v1.HGETRequest) (*v1.HGETResponse, error) {
	resp, err := request.GetTable().Client().GetItemRequest(&dynamodb.GetItemInput{
		ConsistentRead: aws.Bool(true),
		Key: map[string]dynamodb.AttributeValue{
			"pk": {
				S: aws.String(request.GetKey()),
			},
			"sk": {
				S: aws.String(request.GetField()),
			},
		},
		ReturnConsumedCapacity: dynamodb.ReturnConsumedCapacityTotal,
		TableName:              aws.String(request.GetTable().GetTable()),
	}).Send(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.HGETResponse{
		Found: len(resp.Item) > 0,
		Value: &any.Any{
			TypeUrl: aws.StringValue(resp.Item["type"].S),
			Value:   resp.Item["data"].B,
		},
	}, nil
}

func (r RedimoService) HSET(ctx context.Context, request *v1.HSETRequest) (*v1.HSETResponse, error) {
	_, err := request.GetTable().Client().PutItemRequest(&dynamodb.PutItemInput{
		Item: map[string]dynamodb.AttributeValue{
			"pk": {
				S: aws.String(request.GetKey()),
			},
			"sk": {
				S: aws.String(request.GetField()),
			},
			"type": {
				S: aws.String(request.GetValue().GetTypeUrl()),
			},
			"data": {
				B: request.GetValue().GetValue(),
			},
		},
		ReturnConsumedCapacity: dynamodb.ReturnConsumedCapacityTotal,
		ReturnValues:           dynamodb.ReturnValueNone,
		TableName:              aws.String(request.GetTable().GetTable()),
	}).Send(ctx)
	if err != nil {
		return nil, err
	}

	return &v1.HSETResponse{}, nil
}