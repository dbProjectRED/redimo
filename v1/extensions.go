package v1

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/defaults"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func (m Table) Client() *dynamodb.Client {
	config := defaults.Config()
	config.Region = m.GetAwsRegion()
	config.Credentials = aws.NewStaticCredentialsProvider(
		m.GetAwsAccessKeyId(), m.GetAwsSecretAccessKey(), m.GetAwsSessionToken())

	if m.GetEndpointOverride() != "" {
		config.EndpointResolver = aws.ResolveWithEndpointURL(m.GetEndpointOverride())
		config.LogLevel = aws.LogDebugWithHTTPBody
	}

	return dynamodb.New(config)
}
