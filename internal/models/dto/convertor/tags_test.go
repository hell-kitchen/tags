package convertor

import (
	"github.com/google/uuid"
	"github.com/hell-kitchen/tags/internal/models/dto"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

var (
	TestTagCreation1 = &pb.TagCreation{
		Name:  "Дизайн",
		Color: "#0000FF",
		Slug:  "design",
	}
	TestTagCreation2 = &pb.TagCreation{
		Name:  "Программирование",
		Color: "#FF0000",
		Slug:  "programming",
	}
	TestTagCreation3 = &pb.TagCreation{
		Name:  "Маркетинг",
		Color: "#FFA500",
		Slug:  "marketing",
	}
	TestTagCreation4 = &pb.TagCreation{
		Name:  "Реклама",
		Color: "#FF6600",
		Slug:  "advertising",
	}
	TestTagCreation5 = &pb.TagCreation{
		Name:  "Бизнес",
		Color: "#8B008B",
		Slug:  "business",
	}
	TestTagCreation6 = &pb.TagCreation{
		Name:  "Финансы",
		Color: "#0066FF",
		Slug:  "finance",
	}
	TestTagCreation7 = &pb.TagCreation{
		Name:  "Технологии",
		Color: "#993300",
		Slug:  "technology",
	}
	TestTagCreation8 = &pb.TagCreation{
		Name:  "Инновации",
		Color: "#6699FF",
		Slug:  "innovation",
	}
	TestTagCreation9 = &pb.TagCreation{
		Name:  "Образование",
		Color: "#B87333",
		Slug:  "education",
	}
	TestTagCreation10 = &pb.TagCreation{
		Name:  "Менеджмент",
		Color: "#CC9900",
		Slug:  "design",
	}
	TestTag1 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Дизайн",
		Color: "#0000FF",
		Slug:  "design",
	}
	TestTag2 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Программирование",
		Color: "#FF0000",
		Slug:  "programming",
	}
	TestTag3 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Маркетинг",
		Color: "#FFA500",
		Slug:  "marketing",
	}
	TestTag4 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Реклама",
		Color: "#FF6600",
		Slug:  "advertising",
	}
	TestTag5 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Бизнес",
		Color: "#8B008B",
		Slug:  "business",
	}
	TestTag6 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Финансы",
		Color: "#0066FF",
		Slug:  "finance",
	}
	TestTag7 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Технологии",
		Color: "#993300",
		Slug:  "technology",
	}
	TestTag8 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Инновации",
		Color: "#6699FF",
		Slug:  "innovation",
	}
	TestTag9 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Образование",
		Color: "#B87333",
		Slug:  "education",
	}
	TestTag10 = dto.TagDTO{
		ID:    uuid.New(),
		Name:  "Менеджмент",
		Color: "#CC9900",
		Slug:  "design",
	}
)

func TestFromTagCreateRequest_nilData(t *testing.T) {
	t.Run("nil data", func(t *testing.T) {
		assert.Panics(t, func() {
			FromTagCreateRequest(nil)
		})
	})
	t.Run("non initialized object of protobuf type", func(t *testing.T) {
		assert.NotPanics(t, func() {
			FromTagCreateRequest((*pb.CreateRequest)(nil))
		})
	})
	t.Run("main ok case", func(t *testing.T) {
		data := &pb.CreateRequest{
			Name:  "some name",
			Color: "some color",
			Slug:  "some slug",
		}
		got := FromTagCreateRequest(data)
		assert.Equal(t, data.Name, got.Name)
		assert.Equal(t, data.Color, got.Color)
		assert.Equal(t, data.Slug, got.Slug)
	})
}

func TestProtoToUpdateDTO(t *testing.T) {
	req := &pb.UpdateRequest{
		Id:    uuid.NewString(),
		Name:  nil,
		Color: nil,
		Slug:  nil,
	}
	got := ProtoToUpdateDTO(req)
	assert.Equal(t, req.Id, got.ID)
	assert.Equal(t, req.Name, got.Name)
	assert.Equal(t, req.Color, got.Color)
	assert.Equal(t, req.Slug, got.Slug)
}

func TestFromTagCreateManyRequest(t *testing.T) {

	req := &pb.CreateManyRequest{Tags: []*pb.TagCreation{
		TestTagCreation1,
		TestTagCreation2,
		TestTagCreation3,
		TestTagCreation4,
		TestTagCreation5,
		TestTagCreation6,
		TestTagCreation7,
		TestTagCreation8,
		TestTagCreation9,
		TestTagCreation10,
	}}
	resp := FromTagCreateManyRequest(req)
	require.NotNil(t, resp)
	assert.Equal(t, len(req.Tags), len(resp))
}

func TestToTagsCreateManyResponse(t *testing.T) {
	created := []dto.TagDTO{
		TestTag1,
		TestTag2,
		TestTag3,
		TestTag4,
		TestTag5,
		TestTag6,
		TestTag7,
		TestTag8,
		TestTag9,
		TestTag10,
	}
	resp := ToTagsCreateManyResponse(created)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Tags)
	require.Equal(t, len(created), len(resp.Tags))
	for i, tag := range created {
		protoTag := resp.Tags[i]
		if assert.NotNil(t, protoTag) {
			assert.Equal(t, tag.ID.String(), protoTag.Id)
			assert.Equal(t, tag.Name, protoTag.Name)
			assert.Equal(t, tag.Slug, protoTag.Slug)
			assert.Equal(t, tag.Color, protoTag.Color)
		}
	}
}

func TestToTagsGetAllResponse(t *testing.T) {
	all := []dto.TagDTO{
		TestTag1,
		TestTag2,
		TestTag3,
		TestTag4,
		TestTag5,
		TestTag6,
		TestTag7,
		TestTag8,
		TestTag9,
		TestTag10,
	}
	resp := ToTagsGetAllResponse(all)
	require.NotNil(t, resp)
	require.NotNil(t, resp.Tag)
	require.Equal(t, len(all), len(resp.Tag))
	for i, tag := range all {
		protoTag := resp.Tag[i]
		if assert.NotNil(t, protoTag) {
			assert.Equal(t, tag.ID.String(), protoTag.Id)
			assert.Equal(t, tag.Name, protoTag.Name)
			assert.Equal(t, tag.Slug, protoTag.Slug)
			assert.Equal(t, tag.Color, protoTag.Color)
		}
	}
}
