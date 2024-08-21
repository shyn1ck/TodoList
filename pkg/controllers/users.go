package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"todoList/logger"
	"todoList/models"
	"todoList/pkg/repository"
	"todoList/pkg/service"
	"todoList/utils"
)

func GetAllUsers(c *gin.Context) {
	clientIP := c.ClientIP()
	logger.Info.Printf("[controllers.GetAllUsers]: Received request to get all users from IP: %s", clientIP)

	users, err := service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info.Printf("[controllers.GetAllUsers]: Users retrieved successfully from IP: %s", clientIP)
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	clientIP := c.ClientIP()
	logger.Info.Printf("[controllers.GetUserByID]: Received request to get user by ID from IP: %s", clientIP)

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := service.GetUserById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info.Printf("[controllers.GetUserByID]: User retrieved successfully from IP: %s, ID: %v", clientIP, id)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	clientIP := c.ClientIP()
	logger.Info.Printf("[controllers.CreateUser]: Received request to create a new user from IP: %s", clientIP)

	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Info.Printf("[controllers.CreateUser]: User created successfully from IP: %s", clientIP)
	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

func UpdateUser(c *gin.Context) {
	clientIP := c.ClientIP()
	logger.Info.Printf("[controllers.UpdateUser]: Received request to update a user from IP: %s", clientIP)

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	existingUser, err := repository.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if user.Username != "" {
		existingUser.Username = user.Username
	}
	if user.Password != "" {
		existingUser.Password = utils.GenerateHash(user.Password)
	}

	err = repository.UpdateUser(existingUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Printf("[controllers.UpdateUser]: User updated successfully from IP: %s, ID: %v", clientIP, id)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	clientIP := c.ClientIP()
	logger.Info.Printf("[controllers.DeleteUser]: Received request to delete a user from IP: %s", clientIP)

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = service.DeleteUser(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	logger.Info.Printf("[controllers.DeleteUser]: User deleted successfully from IP: %s, ID: %v", clientIP, id)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
