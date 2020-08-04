package main

import (
	"context"

	v1 "github.com/dbProjectRED/redimo/v1"
)

const fieldKey = "/"

func (r RedimoService) Get(ctx context.Context, request *v1.GetRequest) (*v1.GetResponse, error) {
	resp, err := r.HGet(ctx, &v1.HGetRequest{
		Table:     request.GetTable(),
		Key:       request.GetKey(),
		FieldName: fieldKey,
	})
	if err != nil {
		return nil, err
	}

	return &v1.GetResponse{
		Found: resp.GetFound(),
		Value: resp.GetValue(),
	}, nil
}

func (r RedimoService) Set(ctx context.Context, request *v1.SetRequest) (*v1.SetResponse, error) {
	_, err := r.HSet(ctx, &v1.HSetRequest{
		Table:     request.GetTable(),
		Key:       request.GetKey(),
		FieldName: fieldKey,
		Value:     request.GetValue(),
	})

	return &v1.SetResponse{}, err
}

func (r RedimoService) Del(ctx context.Context, request *v1.DelRequest) (*v1.DelResponse, error) {
	_, err := r.HDel(ctx, &v1.HDelRequest{
		Table:     request.GetTable(),
		Key:       request.GetKey(),
		FieldName: fieldKey,
	})

	return &v1.DelResponse{}, err
}
