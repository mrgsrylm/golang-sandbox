package http

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mrgsrylm/simpleshop/pkg/validation"

	dbMocks "github.com/mrgsrylm/simpleshop/pkg/dbs/mocks"
	redisMocks "github.com/mrgsrylm/simpleshop/pkg/redis/mocks"
)

func TestRoutes(t *testing.T) {
	mockDB := dbMocks.NewIDatabase(t)
	mockRedis := redisMocks.NewIRedis(t)
	Routes(gin.New().Group("/"), mockDB, validation.New(), mockRedis)
}
