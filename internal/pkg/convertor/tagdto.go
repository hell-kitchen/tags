package convertor

import (
	"github.com/hell-kitchen/tags/internal/models/dto"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
)

type TagDTOConvertor dto.TagDTO

func FromTagDTO(dto dto.TagDTO) TagDTOConvertor {
	return (TagDTOConvertor)(dto)
}

func (conv TagDTOConvertor) ToProtoTag() *pb.Tag {
	return &pb.Tag{
		Id:    conv.ID.String(),
		Name:  conv.Name,
		Color: conv.Color,
		Slug:  conv.Slug,
	}
}

func (conv TagDTOConvertor) ToProtoUpdateResponse() *pb.UpdateResponse {
	return &pb.UpdateResponse{
		Id:    conv.ID.String(),
		Name:  conv.Name,
		Color: conv.Color,
		Slug:  conv.Slug,
	}
}

func (conv TagDTOConvertor) ToGetResponse() *pb.GetResponse {
	return &pb.GetResponse{
		Tag: conv.ToProtoTag(),
	}
}

func (conv TagDTOConvertor) ToCreateResponse() *pb.CreateResponse {
	return &pb.CreateResponse{
		Id:    conv.ID.String(),
		Name:  conv.Name,
		Color: conv.Color,
		Slug:  conv.Slug,
	}
}
