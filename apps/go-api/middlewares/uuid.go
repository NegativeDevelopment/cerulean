package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CheckUUIDMiddleware(parameterName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		param := c.Param(parameterName)

		paramId, err := uuid.Parse(param)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"message": fmt.Sprintf("Invalid %s given as parameter", parameterName)})
			return
		}

		c.Set(parameterName, paramId)
	}
}
