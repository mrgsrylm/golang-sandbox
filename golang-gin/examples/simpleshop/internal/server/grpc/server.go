package grpc

import (
	"fmt"
	"net"

	"github.com/mrgsrylm/simpleshop/pkg/config"
	"github.com/mrgsrylm/simpleshop/pkg/dbs"
	"github.com/mrgsrylm/simpleshop/pkg/logger"
	"github.com/mrgsrylm/simpleshop/pkg/middleware"
	"github.com/mrgsrylm/simpleshop/pkg/redis"
	"github.com/mrgsrylm/simpleshop/pkg/validation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	cartGRPC "github.com/mrgsrylm/simpleshop/internal/cart/port/grpc"
	userGRPC "github.com/mrgsrylm/simpleshop/internal/user/port/grpc"
)

type Server struct {
	engine    *grpc.Server
	cfg       *config.Schema
	validator validation.Validation
	db        dbs.IDatabase
	cache     redis.IRedis
}

func NewServer(validator validation.Validation, db dbs.IDatabase, cache redis.IRedis) *Server {
	interceptor := middleware.NewAuthInterceptor(config.AuthIgnoreMethods)

	grpcServer := grpc.NewServer(grpc.
		ChainUnaryInterceptor(interceptor.Unary()))

	return &Server{
		engine:    grpcServer,
		cfg:       config.GetConfig(),
		validator: validator,
		db:        db,
		cache:     cache,
	}
}

func (s Server) Run() error {
	userGRPC.RegisterHandlers(s.engine, s.db, s.validator)
	cartGRPC.RegisterHandlers(s.engine, s.db, s.validator)

	reflection.Register(s.engine)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.cfg.GrpcPort))
	logger.Info("GRPC server is listening on PORT: ", s.cfg.GrpcPort)
	if err != nil {
		logger.Error("Failed to listen: ", err)
		return err
	}

	err = s.engine.Serve(lis)
	if err != nil {
		logger.Fatal("Failed to serve grpc: ", err)
		return err
	}

	return nil
}
