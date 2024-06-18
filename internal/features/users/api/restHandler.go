package api

import "github.com/gin-gonic/gin"

type RESTHandler struct {
	service IService
}

func NewREST(service IService) *RESTHandler {
	return &RESTHandler{service: service}
}

func (rest *RESTHandler) Welcome(c *gin.Context) {
	c.JSON(
		200,
		gin.H{"message": "Welcome to Users API"},
	)
}

func (rest *RESTHandler) ListUsers(c *gin.Context) {
	users, err := rest.service.GetUsers()

	if err != nil {
		c.JSON(
			500,
			gin.H{"error": err.Error()},
		)
		return
	}

	c.JSON(
		200,
		gin.H{"message": users},
	)
}
