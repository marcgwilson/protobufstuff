package valuescache

import (
	"google.golang.org/grpc"

	"context"
	"time"
)

type ValuesCacheClient struct {
	Conn   *grpc.ClientConn
	Client *ValuesCacheServiceClient
}

func (this *ValuesCacheClient) Create(cache *ValuesCacheList) (*ValuesCacheResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := *this.Client
	return c.Create(ctx, cache)
}

func (this *ValuesCacheClient) List(resourceID int64) (*ValuesCacheList, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := *this.Client
	return c.List(ctx, &ResourceId{Id: resourceID})
}

func NewValuesCacheClient(conn *grpc.ClientConn) *ValuesCacheClient {
	c := NewValuesCacheServiceClient(conn)
	return &ValuesCacheClient{conn, &c}
}
