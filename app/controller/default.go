package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DefaultController struct {
}

func (l DefaultController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
