package resolver

import (
	hello_pb "apis/go/hello"
	"context"
)

func (r *queryProtoResolver) Hello(ctx context.Context) (*hello_pb.Hello, error) {
	panic("not implemented")
}

type queryProtoResolver struct{ *Resolver }

