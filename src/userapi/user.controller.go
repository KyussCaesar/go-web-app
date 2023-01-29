package userapi

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

type UserController struct {
}

func NewUserController() UserController {
	return UserController{}
}

func (u *UserController) ConfigureRoutes(e *gin.Engine) {
	e.PUT("/user", u.CreateUser)
	e.POST("/user/:id", u.UpdateUser)
	e.GET("/user/:id", u.GetUser)
	e.DELETE("/user/:id", u.DeleteUser)
}

func (u *UserController) CreateUser(c *gin.Context) {
	fmt.Println("create user called")
	c.AbortWithStatus(200)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	c.AbortWithStatus(500)
}

func (u *UserController) GetUser(c *gin.Context) {
	userId := c.Param("id")
	c.String(http.StatusOK, "Hello %s", userId)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	c.AbortWithStatus(500)
}