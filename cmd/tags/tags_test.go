package main

import (
	"github.com/stretchr/testify/require"
	"go.uber.org/fx"
	"testing"
)

func TestNewOptions(t *testing.T) {
	err := fx.ValidateApp(NewOptions())
	require.NoError(t, err)
}
