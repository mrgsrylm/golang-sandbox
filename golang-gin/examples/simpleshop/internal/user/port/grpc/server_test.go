package grpc

import (
	"testing"

	"github.com/mrgsrylm/simpleshop/pkg/dbs/mocks"
	"github.com/mrgsrylm/simpleshop/pkg/validation"

	goGRPC "google.golang.org/grpc"
)

func TestRegisterHandlers(t *testing.T) {
	mockDB := mocks.NewIDatabase(t)
	RegisterHandlers(goGRPC.NewServer(), mockDB, validation.New())
}
