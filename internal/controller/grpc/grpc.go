package grpc

import (
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
	"google.golang.org/grpc"
)

type Controller struct {
	pb.UnimplementedTagsServiceServer
	srv *grpc.Server
}

func New() (*Controller, error) {
	server := grpc.NewServer()
	return &Controller{
		srv: server,
	}, nil
}
