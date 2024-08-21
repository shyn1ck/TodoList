package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoList/models"
	"todoList/pkg/service"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})
}

func SignIn(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, err := service.SignIn(user.Username, user.Password)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{"access_token": accessToken})
}
