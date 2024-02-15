package grpc

import (
	pb "github.com/hell-kitchen/proto/pkg/proto/api/tags"
	"google.golang.org/grpc"
)

type Controller struct {
	pb.UnimplementedTagsServiceServer
}

func New() (*Controller, error) {
	server := grpc.NewServer(grpc.WithTransportCredentials(credentials.J))
}
