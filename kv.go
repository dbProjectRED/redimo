package main

import (
	"context"

	v1 "github.com/dbProjectRED/redimo/v1"
)

func (r RedimoService) GET(ctx context.Context, request *v1.GETRequest) (*v1.GETResponse, error) {
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

func (r RedimoService) SET(ctx context.Context, request *v1.SETRequest) (*v1.SETResponse, error) {
	_, err := r.HSET(ctx, &v1.HSETRequest{
		Table: request.GetTable(),
		Key:   request.GetKey(),
		Field: "/",
		Value: request.GetValue(),
	})

	return &v1.SETResponse{}, err
}
