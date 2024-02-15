package grpc

import (
	"errors"
	"fmt"
	"github.com/hell-kitchen/tags/internal/config"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
	"go.uber.org/multierr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net"
)

type Controller struct {
	pb.UnimplementedTagsServiceServer
	server   *grpc.Server
	cfg      *config.Config
	listener net.Listener
}

func New(cfg *config.Config) (ctrl *Controller, err error) {
	ctrl = &Controller{
		cfg: cfg,
	}
	multierr.AppendInto(&err, ctrl.createListener())
	multierr.AppendInto(&err, ctrl.createServer())
	return ctrl, err
}

func (ctrl *Controller) createListener() (err error) {
	ctrl.listener, err = net.Listen("tcp", ctrl.cfg.Bind())
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
