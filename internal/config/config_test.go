package config

import (
	"context"
	"errors"
	"github.com/hell-kitchen/pkg/confita"
	"github.com/stretchr/testify/assert"
	"testing"
)

var errUnknown = errors.New("")

type loaderMock struct {
	f func(context.Context, any) error
}

func (mock loaderMock) Load(ctx context.Context, to any) error {
	return mock.f(ctx, to)
}

func prepareLoader(t testing.TB, f func(ctx context.Context, to interface{}) error) func() {
	t.Helper()

	_ = confita.Get()
	var loader confita.Loader = loaderMock{f: f}
	return confita.SetNewLoader(loader)
}

func TestNewController_OK(t *testing.T) {
	var calledTimes int
	f := func(ctx context.Context, to interface{}) error {
		calledTimes++
		return nil
	}

	defer prepareLoader(t, f)()

	ctrl, err := NewController()
	assert.NoError(t, err)
	assert.NotEmpty(t, calledTimes)
	if assert.NotNil(t, ctrl) {
		assert.Equal(t, Controller{
			BindPort: 8080,
			BindHost: "0.0.0.0",
			BaseAddr: "http://localhost:8080",
			UseTLS:   false,
			CertFile: "",
			KeyFile:  "",
		}, *ctrl)
	}
}

func TestNewController_NonNilError(t *testing.T) {
	var calledTimes int

	f := func(ctx context.Context, to interface{}) error {
		calledTimes++
		return errUnknown
	}

	defer prepareLoader(t, f)()

	ctrl, err := NewController()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errUnknown)
	}
	if assert.NotEmpty(t, calledTimes) {
		assert.Equal(t, 1, calledTimes)
	}
	assert.Nil(t, ctrl)
}

func TestController_Bind(t *testing.T) {
	ctrl := Controller{
		BindPort: 123,
		BindHost: "localhost",
	}
	res := ctrl.Bind()
	assert.Equal(t, "localhost:123", res)
}
