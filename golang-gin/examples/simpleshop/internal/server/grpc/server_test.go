package grpc

import (
	"testing"

	"github.com/mrgsrylm/simpleshop/pkg/validation"
	"github.com/stretchr/testify/assert"

	dbMocks "github.com/mrgsrylm/simpleshop/pkg/dbs/mocks"
	redisMocks "github.com/mrgsrylm/simpleshop/pkg/redis/mocks"
)

func TestNewServer(t *testing.T) {
	mockDB := dbMocks.NewIDatabase(t)
	mockRedis := redisMocks.NewIRedis(t)

	server := NewServer(validation.New(), mockDB, mockRedis)
	assert.NotNil(t, server)
}
