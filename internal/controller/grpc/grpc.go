package grpc

import (
	"context"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/hell-kitchen/pkg/grpcmw"
	"github.com/hell-kitchen/pkg/logger"
	"github.com/hell-kitchen/tags/internal/config"
	"github.com/hell-kitchen/tags/internal/service"
	pb "github.com/hell-kitchen/tags/pkg/api/proto"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"time"
)

type Controller struct {
	pb.UnimplementedTagsServiceServer
	log      *zap.Logger
	server   *grpc.Server
	cfg      *config.Controller
	listener net.Listener
	service  service.TagsService
}

func New(cfg *config.Controller, service service.TagsService, log *zap.Logger) (ctrl *Controller, err error) {
	ctrlLogger := log.With(
		logger.WithLayer("controller"),
		zap.String("bind-addr", cfg.Bind()),
	)
	ctrl = &Controller{
		cfg:     cfg,
		log:     ctrlLogger,
		service: service,
	}
	multierr.AppendInto(&err, ctrl.createListener())
	multierr.AppendInto(&err, ctrl.createServer(log))
	return ctrl, err
}

func (ctrl *Controller) createListener() (err error) {
	ctrl.listener, err = net.Listen("tcp", ctrl.cfg.Bind())
	if err != nil {
		return err
	}
	ctrl.log.Info("created listener")
	return nil
}

func transportCredentialsByConfig(cfg *config.Controller) credentials.TransportCredentials {
	if !cfg.UseTLS {
		return insecure.NewCredentials()
	}

	if cfg.CertFile == "" || cfg.KeyFile == "" {
		return insecure.NewCredentials()
	}

	creds, err := credentials.NewServerTLSFromFile(cfg.CertFile, cfg.KeyFile)
	if err != nil {
		return insecure.NewCredentials()
	}

	return creds
}

func (ctrl *Controller) createServer(log *zap.Logger) (err error) {
	grpcLogger := log.With(logger.WithLayer("grpc"))

	grpc_zap.ReplaceGrpcLoggerV2(grpcLogger)

	ctrl.log.Info("creating gRPC server")
	var opts = []grpc.ServerOption{
		grpc.Creds(transportCredentialsByConfig(ctrl.cfg)),
		grpc.ChainUnaryInterceptor(
			grpc_zap.UnaryServerInterceptor(grpcLogger),
			grpc_recovery.UnaryServerInterceptor(),
			grpcmw.UnaryServerInterceptor(),
		),
	}

	ctrl.server = grpc.NewServer(opts...)
	pb.RegisterTagsServiceServer(ctrl.server, ctrl)
	ctrl.log.Info("successfully created gRPC server")
	return nil
}

func (ctrl *Controller) Start(ctx context.Context) error {
	var cancel context.CancelCauseFunc
	ctrl.log.Info("Start just called")

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
