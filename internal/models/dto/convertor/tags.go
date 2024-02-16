package convertor

import (
	"github.com/hell-kitchen/tags/internal/models/dto"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
)

func FromTagCreateRequest(pb *pb.TagsCreateRequest) dto.TagCreationDTO {
	return dto.TagCreationDTO{
		Name:  pb.GetName(),
		Color: pb.GetColor(),
		Slug:  pb.GetSlug(),
	}
}

func ToTagsCreateResponse(dto *dto.TagDTO) *pb.TagsCreateResponse {
	if dto == nil {
		return nil
	}
	return &pb.TagsCreateResponse{
		Id:    dto.ID.String(),
		Name:  dto.Name,
		Color: dto.Color,
		Slug:  dto.Slug,
	}
}

func ToTagsGetResponse(dto *dto.TagDTO) *pb.TagsGetResponse {
	if dto == nil {
		return nil
	}
	return &pb.TagsGetResponse{
		Tag: &pb.Tag{
			Id:    dto.ID.String(),
			Name:  dto.Name,
			Color: dto.Color,
			Slug:  dto.Slug,
		},
	}
}

func ToTagsGetAllResponse(tags []dto.TagDTO) *pb.TagsGetAllResponse {
	var result = make([]*pb.Tag, 0, len(tags))
	for _, tag := range tags {
		result = append(result, &pb.Tag{
			Id:    tag.ID.String(),
			Name:  tag.Name,
			Color: tag.Color,
			Slug:  tag.Slug,
		})
	}
	return &pb.TagsGetAllResponse{Tag: result}
}
