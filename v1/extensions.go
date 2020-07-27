package v1

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (t Table) Client() *dynamodb.Client {
	return dynamodb.New(aws.Config{
		Region:           "",
		Credentials:      aws.NewStaticCredentialsProvider("", "", ""),
		EndpointResolver: endpoints.NewDefaultResolver(),
	})
}
