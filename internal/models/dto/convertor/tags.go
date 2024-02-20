package convertor

import (
	"github.com/hell-kitchen/tags/internal/models/dto"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
)

type CreationRequest interface {
	GetName() string
	GetColor() string
	GetSlug() string
}

func FromTagCreateRequest(pb CreationRequest) dto.TagCreationDTO {
	return dto.TagCreationDTO{
		Name:  pb.GetName(),
		Color: pb.GetColor(),
		Slug:  pb.GetSlug(),
	}
}

func FromTagCreateManyRequest(pb *pb.CreateManyRequest) []dto.TagCreationDTO {
	result := make([]dto.TagCreationDTO, 0, len(pb.GetTags()))
	for _, i := range pb.GetTags() {
		result = append(result, FromTagCreateRequest(i))
	}
	return result
}

func ToTagsCreateManyResponse(created []dto.TagDTO) *pb.CreateManyResponse {
	resp := &pb.CreateManyResponse{
		Tags: make([]*pb.Tag, 0, len(created)),
	}
	for _, tag := range created {
		resp.Tags = append(resp.Tags, &pb.Tag{
			Id:    tag.ID.String(),
			Name:  tag.Name,
			Color: tag.Color,
			Slug:  tag.Slug,
		})
	}
	return resp
}

func ProtoToUpdateDTO(proto *pb.UpdateRequest) dto.TagUpdateDTO {
	return dto.TagUpdateDTO{
		ID:    proto.Id,
		Name:  proto.Name,
		Color: proto.Color,
		Slug:  proto.Slug,
	}
}

func ToTagsGetAllResponse(tags []dto.TagDTO) *pb.GetAllResponse {
	var result = make([]*pb.Tag, 0, len(tags))
	for _, tag := range tags {
		result = append(result, &pb.Tag{
			Id:    tag.ID.String(),
			Name:  tag.Name,
			Color: tag.Color,
			Slug:  tag.Slug,
		})
	}
	return &pb.GetAllResponse{Tag: result}
}
