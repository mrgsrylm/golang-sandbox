package http

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrgsrylm/simpleshop/pkg/dbs/mocks"
	"github.com/mrgsrylm/simpleshop/pkg/validation"
)

func TestRoutes(t *testing.T) {
	mockDB := mocks.NewIDatabase(t)
	Routes(gin.New().Group("/"), mockDB, validation.New())
}
