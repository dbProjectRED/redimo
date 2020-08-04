package main_test

import (
	"context"
	"testing"

	redimo "github.com/dbProjectRED/redimo"
	v1 "github.com/dbProjectRED/redimo/v1"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/wrappers"
	"github.com/stretchr/testify/assert"
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
	assert.True(t, resp.Found)

	sv := wrappers.StringValue{}
	assert.NoError(t, ptypes.UnmarshalAny(resp.Value, &sv))
	assert.Equal(t, "v1", sv.Value)

	_, err = service.Del(context.Background(), &v1.DelRequest{
		Table: &table,
		Key:   "k1",
	})
	assert.NoError(t, err)

	resp, err = service.Get(context.Background(), &v1.GetRequest{
		Table: &table,
		Key:   "k1",
	})
	assert.NoError(t, err)
	assert.False(t, resp.Found)
}
