package grpc

import (
	"context"
	"github.com/hell-kitchen/pkg/grpcmw"
	old "github.com/hell-kitchen/tags/internal/models/dto/convertor"
	"github.com/hell-kitchen/tags/internal/pkg/convertor"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
	"go.uber.org/zap"
)

func requestIDField(ctx context.Context) zap.Field {
	requestID := grpcmw.FromContext(ctx)
	return zap.String("request-id", requestID)
}

func (ctrl *Controller) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	requestID := requestIDField(ctx)
	ctrl.log.Debug("got create request", requestID)

	resp, err := ctrl.service.Create(ctx, old.FromTagCreateRequest(req), requestID)
	if err != nil {
		return nil, err
	}
	return convertor.FromTagDTO(*resp).ToCreateResponse(), nil
}

func (ctrl *Controller) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	requestID := requestIDField(ctx)
	ctrl.log.Debug("got Get by id request",
		requestID,
		zap.String("tag-id", req.GetId()),
	)

	resp, err := ctrl.service.Get(ctx, req.GetId(), requestID)
	if err != nil {
		return nil, err
	}

	return convertor.FromTagDTO(*resp).ToGetResponse(), nil
}

func (ctrl *Controller) GetAll(ctx context.Context, _ *pb.GetAllRequest) (*pb.GetAllResponse, error) {
	requestID := requestIDField(ctx)
	ctrl.log.Debug("got Get all request",
		requestID,
	)

	resp, err := ctrl.service.GetAll(ctx, requestID)
	if err != nil {
		return nil, err
	}
	return old.ToTagsGetAllResponse(resp), nil
}

func (ctrl *Controller) CreateMany(ctx context.Context, req *pb.CreateManyRequest) (*pb.CreateManyResponse, error) {
	requestID := requestIDField(ctx)
	ctrl.log.Debug("got CreateMany request",
		requestID,
	)

	resp, err := ctrl.service.CreateMany(ctx, old.FromTagCreateManyRequest(req), requestID)
	if err != nil {
		ctrl.log.Debug("got non nil error", requestID, zap.Error(err))
		return nil, err
	}
	return old.ToTagsCreateManyResponse(resp), nil
}

func (ctrl *Controller) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	requestID := requestIDField(ctx)
	ctrl.log.Debug("got Get all request",
		requestID,
	)

	err := ctrl.service.Delete(ctx, req.GetId(), requestID)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{}, nil
}

func (ctrl *Controller) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	requestID := requestIDField(ctx)
	ctrl.log.Debug("got update request", requestID)

	resp, err := ctrl.service.Update(ctx, old.ProtoToUpdateDTO(req), requestID)
	if err != nil {
		return nil, err
	}
	return convertor.FromTagDTO(*resp).ToProtoUpdateResponse(), nil
}
