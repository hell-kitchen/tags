package grpc

import (
	"context"
	"github.com/hell-kitchen/tags/internal/models/dto/convertor"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
)

func (ctrl *Controller) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	resp, err := ctrl.service.Create(ctx, convertor.FromTagCreateRequest(req))
	if err != nil {
		return nil, err
	}
	return convertor.ToTagsCreateResponse(resp), nil
}

func (ctrl *Controller) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	resp, err := ctrl.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return convertor.ToTagsGetResponse(resp), nil
}

func (ctrl *Controller) GetAll(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	resp, err := ctrl.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return convertor.ToTagsGetAllResponse(resp), nil
}

func (ctrl *Controller) CreateMany(context.Context, *pb.CreateManyRequest) (*pb.CreateManyResponse, error) {
	panic("not implemented")
}

func (ctrl *Controller) Delete(context.Context, *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	panic("not implemented")
}
func (ctrl *Controller) Update(context.Context, *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	panic("not implemented")
}
