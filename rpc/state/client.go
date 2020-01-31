package state

import (
	"github.com/marcgwilson/protobufstuff/mode"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"context"
	"time"
)

type CacheStateClient struct {
	Conn   *grpc.ClientConn
	Client *CacheStateServiceClient
}

func (this *CacheStateClient) SetCacheState(resource string, m mode.CacheMode) (*CacheStateResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := *this.Client
	return c.SetCacheState(ctx, &CacheState{Resource: resource, Mode: uint32(m)})
}

func (this *CacheStateClient) GetCacheState() (*CacheState, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := *this.Client
	return c.GetCacheState(ctx, &empty.Empty{})
}

func NewCacheStateClient(conn *grpc.ClientConn) *CacheStateClient {
	c := NewCacheStateServiceClient(conn)
	return &CacheStateClient{conn, &c}
}
