package grpc

import (
	"context"
	"errors"
	"fmt"
	"github.com/hell-kitchen/tags/internal/config"
	"github.com/hell-kitchen/tags/internal/service"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
	"time"
)

var listenerFunc func(network string, address string) (net.Listener, error) = net.Listen

type Controller struct {
	pb.UnimplementedTagsServiceServer
	server   *grpc.Server
	cfg      *config.Controller
	listener net.Listener
	service  service.TagsService
}

func New(cfg *config.Controller, service service.TagsService) (ctrl *Controller, err error) {
	ctrl = &Controller{
		cfg:     cfg,
		service: service,
	}
	multierr.AppendInto(&err, ctrl.createListener())
	multierr.AppendInto(&err, ctrl.createServer())
	return ctrl, err
}

func (ctrl *Controller) createListener() (err error) {
	ctrl.listener, err = listenerFunc("tcp", ctrl.cfg.Bind())
	if err != nil {
		return err
	}
	return nil
}

func (ctrl *Controller) createServer() (err error) {
	var opts []grpc.ServerOption

	if ctrl.cfg.UseTLS {
		var transportCredentials credentials.TransportCredentials

		if ctrl.cfg.CertFile == "" || ctrl.cfg.KeyFile == "" {
			return errors.New("bad config")
		}

		transportCredentials, err = credentials.NewServerTLSFromFile(ctrl.cfg.CertFile, ctrl.cfg.KeyFile)
		if err != nil {
			return fmt.Errorf("failed to generate credentials: %w", err)
		}

		opts = []grpc.ServerOption{
			grpc.Creds(transportCredentials),
		}
	}

	ctrl.server = grpc.NewServer(opts...)
	pb.RegisterTagsServiceServer(ctrl.server, ctrl)
	return nil
}

func (ctrl *Controller) Start(ctx context.Context) error {
	var cancel context.CancelCauseFunc

	ctx, cancel = context.WithCancelCause(ctx)

	go func() {
		err := ctrl.server.Serve(ctrl.listener)
		if err != nil {
			cancel(err)
		}
	}()

	time.Sleep(100 * time.Millisecond)
	return ctx.Err()
}

func (ctrl *Controller) Stop(_ context.Context) error {
	ctrl.server.GracefulStop()
	return nil
}
