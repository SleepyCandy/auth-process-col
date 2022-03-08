package authen

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	h := AuthHandler{}
	groupAuth := r.Group("/loc-process-authentication/v1")
	{
		groupAuth.POST("/loanOriginatedAurgentication", h.ServiceAuth)
		groupAuth.GET("/loanOriginatedAurgentication/get", func(context *gin.Context) {
			context.JSON(200, "hello")
		})
	}

	return r
}
