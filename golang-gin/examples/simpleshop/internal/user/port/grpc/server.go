package grpc

import (
	"github.com/mrgsrylm/simpleshop/internal/user/repository"
	"github.com/mrgsrylm/simpleshop/internal/user/service"
	"github.com/mrgsrylm/simpleshop/pkg/dbs"
	"github.com/mrgsrylm/simpleshop/pkg/validation"
	"google.golang.org/grpc"

	pb "github.com/mrgsrylm/simpleshop/proto/gen/go/user"
)

func RegisterHandlers(svr *grpc.Server, db dbs.IDatabase, validator validation.Validation) {
	userRepo := repository.NewUserRepository(db)
	userSvc := service.NewUserService(validator, userRepo)
	userHandler := NewUserHandler(userSvc)

	pb.RegisterUserServiceServer(svr, userHandler)
}
