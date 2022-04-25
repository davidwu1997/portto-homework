package api

import (
	"github.com/gin-gonic/gin"

	"portto-homework/pkg/context"
)

func InjectContext(c *gin.Context) {
	c.Set("ctx", context.Background())
	c.Next()
}
