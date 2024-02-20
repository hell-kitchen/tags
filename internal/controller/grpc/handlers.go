package grpc

import (
	"context"
	old "github.com/hell-kitchen/tags/internal/models/dto/convertor"
	"github.com/hell-kitchen/tags/internal/pkg/convertor"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
)

func (ctrl *Controller) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	resp, err := ctrl.service.Create(ctx, old.FromTagCreateRequest(req))
	if err != nil {
		return nil, err
	}
	return convertor.FromTagDTO(*resp).ToCreateResponse(), nil
}

func (ctrl *Controller) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	resp, err := ctrl.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return convertor.FromTagDTO(*resp).ToGetResponse(), nil
}

func (ctrl *Controller) GetAll(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	resp, err := ctrl.service.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return old.ToTagsGetAllResponse(resp), nil
}

func (ctrl *Controller) CreateMany(ctx context.Context, req *pb.CreateManyRequest) (*pb.CreateManyResponse, error) {
	resp, err := ctrl.service.CreateMany(ctx, old.FromTagCreateManyRequest(req))
	if err != nil {
		return nil, err
	}
	return old.ToTagsCreateManyResponse(resp), nil
}

func (ctrl *Controller) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	err := ctrl.service.Delete(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (ctrl *Controller) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	resp, err := ctrl.service.Update(ctx, old.ProtoToUpdateDTO(req))
	if err != nil {
		return nil, err
	}
	return convertor.FromTagDTO(*resp).ToProtoUpdateResponse(), nil
}
