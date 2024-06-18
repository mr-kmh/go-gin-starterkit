package api

import "github.com/gin-gonic/gin"

func REST(r *gin.Engine, rest *RESTHandler) {
	userRouter := r.Group("/users")
	{
		userRouter.GET("/welcome", rest.Welcome)
		userRouter.GET("", rest.ListUsers)
	}
}
