package config

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewPostgres_OK(t *testing.T) {
	var calledTimes int
	f := func(ctx context.Context, to interface{}) error {
		calledTimes++
		return nil
	}

	defer prepareLoader(t, f)()

	ctrl, err := NewPostgres()
	assert.NoError(t, err)
	assert.NotEmpty(t, calledTimes)
	if assert.NotNil(t, ctrl) {
		assert.Equal(t, &Postgres{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "postgres",
			Database: "postgres",
		}, ctrl)
	}
}

func TestNewPostgres_UnknownError(t *testing.T) {
	var calledTimes int
	errUnknown := errors.New("")
	f := func(ctx context.Context, to interface{}) error {
		calledTimes++
		return errUnknown
	}

	defer prepareLoader(t, f)()

	ctrl, err := NewPostgres()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errUnknown)
	}
	assert.NotEmpty(t, calledTimes)
	assert.Nil(t, ctrl)
}

func TestPostgres_ConnString(t *testing.T) {
	cfg := Postgres{
		Host:     "some host",
		Port:     123123,
		User:     "some user",
		Password: "some password",
		Database: "some db",
	}
	res := cfg.ConnString()
	assert.Equal(t, "postgresql://some user:some password@some host:123123/some db?sslmode=disable", res)
}
