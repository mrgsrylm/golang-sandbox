package response

import (
	"github.com/gin-gonic/gin"

	"github.com/mrgsrylm/simpleshop/pkg/config"
)

func Error(c *gin.Context, status int, err error, message string) {
	cfg := config.GetConfig()
	errorRes := map[string]interface{}{
		"message": message,
	}

	if cfg.Environment != config.ProductionEnv {
		errorRes["debug"] = err.Error()
	}

	c.JSON(status, Response{Error: errorRes})
}
