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

func ToTagsCreateResponse(dto *dto.TagDTO) *pb.CreateResponse {
	if dto == nil {
		return nil
	}
	return &pb.CreateResponse{
		Id:    dto.ID.String(),
		Name:  dto.Name,
		Color: dto.Color,
		Slug:  dto.Slug,
	}
}

func ToTagsGetResponse(dto *dto.TagDTO) *pb.GetResponse {
	if dto == nil {
		return nil
	}
	return &pb.GetResponse{
		Tag: &pb.Tag{
			Id:    dto.ID.String(),
			Name:  dto.Name,
			Color: dto.Color,
			Slug:  dto.Slug,
		},
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
