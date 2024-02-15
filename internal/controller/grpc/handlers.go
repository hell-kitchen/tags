package grpc

import (
	"context"
	"github.com/hell-kitchen/tags/internal/models/dto/convertors"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
)

func (ctrl *Controller) Create(ctx context.Context, req *pb.TagsCreateRequest) (*pb.TagsCreateResponse, error) {
	resp, err := ctrl.service.Create(ctx, convertors.FromTagCreateRequest(req))
	if err != nil {
		return nil, err
	}
	return convertors.ToTagsCreateResponse(resp), nil
}

func (ctrl *Controller) Get(ctx context.Context, req *pb.TagsGetRequest) (*pb.TagsGetResponse, error) {
	resp, err := ctrl.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return convertors.ToTagsGetResponse(resp), nil
}

func (ctrl *Controller) GetAll(ctx context.Context, _ *pb.TagsGetAllRequest) (*pb.TagsGetAllResponse, error) {
	resp, err := ctrl.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return convertors.ToTagsGetAllResponse(resp), nil
}
