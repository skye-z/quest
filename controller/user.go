package controller

import (
	"net/http"
	"quest/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func (uc UserController) GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	user, err := uc.UserService.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
