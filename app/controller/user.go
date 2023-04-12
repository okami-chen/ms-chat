package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func (u UserController) GetUserInfo(c *gin.Context) {
	username := c.Query("username")
	age := c.Query("age")
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"age":      age,
	})
}
