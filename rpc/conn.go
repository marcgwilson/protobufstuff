package rpc

import (
	"github.com/marcgwilson/protobufstuff/config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func NewConn(cfg *config.RPCConfig) (*grpc.ClientConn, error) {
	var opts []grpc.DialOption
	if cfg.TLS != nil {
		creds, err := credentials.NewClientTLSFromFile(cfg.TLS.Cert, "")
		if err != nil {
			return nil, err
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// Return connection to the server.
	return grpc.Dial(cfg.Address(), opts...)
}
