package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"todoList/models"
	"todoList/pkg/repository"
	"todoList/pkg/service"
	"todoList/utils"
)

func GetAllUsers(c *gin.Context) {
	log.Println("controllers.GetAllUsers: Received request to get all users")
	users, err := service.GetAllUsers()
	if err != nil {
		log.Printf("controllers.GetAllUsers: Failed to retrieve users, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("controllers.GetAllUsers: Users retrieved successfully")
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func GetUserByID(c *gin.Context) {
	log.Println("controllers.GetUserByID: Received request to get user by ID")
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("controllers.GetUserByID: Invalid user ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	user, err := service.GetUserById(uint(id))
	if err != nil {
		log.Printf("controllers.GetUserByID: Failed to retrieve user, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Printf("controllers.GetUserByID: User retrieved successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	log.Println("controllers.CreateUser: Received request to create a new user")
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		log.Printf("controllers.CreateUser: Failed to bind JSON, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := service.CreateUser(user)
	if err != nil {
		log.Printf("controllers.CreateUser: Failed to create user, error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	log.Println("controllers.CreateUser: User created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

func UpdateUser(c *gin.Context) {
	log.Println("controllers.UpdateUser: Received request to update a user")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Printf("controllers.UpdateUser: Failed to bind JSON, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		log.Printf("controllers.UpdateUser: Invalid user ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	existingUser, err := repository.GetUserByID(uint(id))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("controllers.UpdateUser: User not found, ID: %v\n", id)
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		log.Printf("controllers.UpdateUser: Failed to retrieve user, ID: %v, error: %v\n", id, err)
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
		log.Printf("controllers.UpdateUser: Failed to update user, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("controllers.UpdateUser: User updated successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func DeleteUser(c *gin.Context) {
	log.Println("controllers.DeleteUser: Received request to delete a user")
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("controllers.DeleteUser: Invalid user ID, error: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	err = service.DeleteUser(uint(id))
	if err != nil {
		log.Printf("controllers.DeleteUser: Failed to delete user, ID: %v, error: %v\n", id, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("controllers.DeleteUser: User deleted successfully, ID: %v\n", id)
	c.JSON(http.StatusOK, gin.H{"message": "user deleted successfully"})
}
