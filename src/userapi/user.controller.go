package userapi

import (
	"fmt"
	"net/http"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/kyusscaesar/go-web-app/client"
)

type UserController struct {
	userService *UserService
}

func NewUserController(u *UserService) UserController {
	return UserController{
		userService: u
	}
}

func (u *UserController) ConfigureRoutes(e *gin.Engine) {
	e.POST("/user", u.CreateUser)
	e.PUT("/user/:id", u.UpdateUser)
	e.GET("/user/:id", u.GetUser)
	e.DELETE("/user/:id", u.DeleteUser)
}

func (u *UserController) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	var dto client.ICreateUserDTO
	err := c.BindJSON(&dto)
	if err != nil {
		c.AbortWithStatus(400)
		return
	}

	user, err := u.userService.CreateUser(ctx, dto.Email, dto.Username)
	if err != nil {
		if errors.Is(err, ErrUserAlreadyExists) {
			c.AbortWithStatus(409)
		} else {
			c.AbortWithStatus(500)
		}

		return
	}

	c.JSON(201, user)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	c.AbortWithStatus(500) // TODO
}

func (u *UserController) GetUser(c *gin.Context) {
	userId := c.Param("id")
	c.String(http.StatusOK, "Hello %s", userId)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	c.AbortWithStatus(500)
}