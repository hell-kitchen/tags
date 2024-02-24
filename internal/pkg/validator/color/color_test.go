package color

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidate(t *testing.T) {
	tt := []struct {
		validate string
		valid    bool
	}{
		{"#FFFFFF", true},
		{"#FFFFFA", true},
		{"#aaaaaa", true},
		{"#1e", false},
		{"#ZZZZZZ", false},
		{"#ZZZZZZ", false},
		{"#000333", true},
	}
	for _, tc := range tt {
		t.Run(tc.validate, func(t *testing.T) {
			assert.Equal(t, tc.valid, Validate(tc.validate))
		})
	}
}
