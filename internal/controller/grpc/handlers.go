package grpc

import (
	"context"
	pb "github.com/hell-kitchen/proto/pkg/proto/api/tags"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (ctrl *Controller) Create(context.Context, *pb.TagsCreateRequest) (*pb.TagsCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}

func (ctrl *Controller) Get(context.Context, *pb.TagsGetRequest) (*pb.TagsGetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}

func (ctrl *Controller) GetAll(context.Context, *pb.TagsGetAllRequest) (*pb.TagsGetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
