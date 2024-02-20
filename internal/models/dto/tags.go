package dto

import "github.com/google/uuid"

type (
	TagCreationDTO struct {
		Name  string
		Color string
		Slug  string
	}
	TagUpdateDTO struct {
		ID    string
		Name  *string
		Color *string
		Slug  *string
	}
	TagDTO struct {
		ID    uuid.UUID
		Name  string
		Color string
		Slug  string
	}
)
