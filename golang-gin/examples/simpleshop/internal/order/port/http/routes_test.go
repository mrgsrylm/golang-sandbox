package http

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrgsrylm/simpleshop/pkg/validation"

	dbMocks "github.com/mrgsrylm/simpleshop/pkg/dbs/mocks"
)

func TestRoutes(t *testing.T) {
	mockDB := dbMocks.NewIDatabase(t)
	Routes(gin.New().Group("/"), mockDB, validation.New())
}
