package main

import (
	"context"
	"log"
	"net"

	"github.com/aws/aws-sdk-go-v2/aws"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"

	v1 "github.com/dbProjectRED/redimo/v1"
)

func main() {
	lis, err := net.Listen("tcp", ":4772")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	v1.RegisterRedimoServiceServer(server, redimoService{})
	log.Fatal(server.Serve(lis))
}

type redimoService struct{}

func (r redimoService) GET(ctx context.Context, request *v1.GETRequest) (*v1.GETResponse, error) {
	resp, err := r.HGET(ctx, &v1.HGETRequest{
		Table: request.GetTable(),
		Key:   request.GetKey(),
		Field: "/",
	})
	if err != nil {
		return nil, err
	}
	return &v1.GETResponse{
		Found: resp.GetFound(),
		Value: resp.GetValue(),
	}, nil
}

func (r redimoService) SET(ctx context.Context, request *v1.SETRequest) (*v1.SETResponse, error) {
	_, err := r.HSET(ctx, &v1.HSETRequest{
		Table: request.GetTable(),
		Key:   request.GetKey(),
		Field: "/",
		Value: request.GetValue(),
	})
	return &v1.SETResponse{}, err
}

func (r redimoService) HGET(ctx context.Context, request *v1.HGETRequest) (*v1.HGETResponse, error) {
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

func (r redimoService) HSET(ctx context.Context, request *v1.HSETRequest) (*v1.HSETResponse, error) {
	_, err := request.GetTable().Client().PutItemRequest(&dynamodb.PutItemInput{
		Item: map[string]dynamodb.AttributeValue{
			"pk": {
				S: aws.String(request.GetField()),
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
