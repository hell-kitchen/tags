package grpc

import (
	"errors"
	"github.com/hell-kitchen/tags/internal/config"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

type myListener struct {
}

func (m myListener) Accept() (net.Conn, error) {
	return nil, nil
}

func (m myListener) Close() error {
	return nil
}

func (m myListener) Addr() net.Addr {
	return nil
}

func TestController_createListener_OK(t *testing.T) {
	listener := myListener{}
	listenerFunc = func(network string, address string) (net.Listener, error) {
		return listener, nil
	}
	ctrl := &Controller{
		cfg: &config.Controller{},
	}
	assert.NotEqual(t, ctrl.listener, listener)
	err := ctrl.createListener()
	assert.NoError(t, err)
	if assert.NotNil(t, ctrl.listener) {
		assert.Equal(t, listener, ctrl.listener)
	}
}

func TestController_createListener_SomeError(t *testing.T) {
	var errUnknown = errors.New("")
	listener := myListener{}
	listenerFunc = func(network string, address string) (net.Listener, error) {
		return nil, errUnknown
	}
	ctrl := &Controller{
		cfg: &config.Controller{},
	}
	assert.NotEqual(t, ctrl.listener, listener)
	err := ctrl.createListener()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errUnknown)
	}
	assert.Equal(t, nil, ctrl.listener)
}
