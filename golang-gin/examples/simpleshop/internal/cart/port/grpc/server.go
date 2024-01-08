package grpc

import (
	"google.golang.org/grpc"

	"github.com/mrgsrylm/simpleshop/internal/cart/repository"
	"github.com/mrgsrylm/simpleshop/internal/cart/service"
	"github.com/mrgsrylm/simpleshop/pkg/dbs"
	"github.com/mrgsrylm/simpleshop/pkg/validation"

	pb "github.com/mrgsrylm/simpleshop/proto/gen/go/cart"
)

func RegisterHandlers(svr *grpc.Server, db dbs.IDatabase, validator validation.Validation) {
	cartRepo := repository.NewCartRepository(db)
	cartSvc := service.NewCartService(validator, cartRepo)
	cartHandler := NewCartHandler(cartSvc)

	pb.RegisterCartServiceServer(svr, cartHandler)
}
