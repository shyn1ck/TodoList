package repository

import (
	"log"
	"todoList/db"
	"todoList/models"
)

func GetAllUsers() (users []models.User, err error) {
	log.Println("repository.GetAllUsers: Fetching all users from the database")
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		log.Printf("repository.GetAllUsers: Failed to fetch users, error: %v\n", err)
		return nil, err
	}

	log.Println("repository.GetAllUsers: Successfully fetched all users")
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	log.Printf("repository.GetUserByID: Fetching user by ID %v from the database\n", id)
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		log.Printf("repository.GetUserByID: Failed to fetch user, ID: %v, error: %v\n", id, err)
		return user, err
	}

	log.Printf("repository.GetUserByID: Successfully fetched user by ID %v\n", id)
	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	log.Printf("repository.GetUserByUsername: Fetching user by username %v from the database\n", username)
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Printf("repository.GetUserByUsername: Failed to fetch user by username %v, error: %v\n", username, err)
		return user, err
	}

	log.Printf("repository.GetUserByUsername: Successfully fetched user by username %v\n", username)
	return user, nil
}

func CreateUser(user models.User) (err error) {
	log.Println("repository.CreateUser: Creating a new user in the database")
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		log.Printf("repository.CreateUser: Failed to create user, error: %v\n", err)
		return err
	}

	log.Println("repository.CreateUser: User created successfully")
	return nil
}

func UpdateUser(user models.User) (err error) {
	log.Printf("repository.UpdateUser: Updating user with ID %v in the database\n", user.ID)
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		log.Printf("repository.UpdateUser: Failed to update user with ID %v, error: %v\n", user.ID, err)
		return err
	}

	log.Printf("repository.UpdateUser: User with ID %v updated successfully\n", user.ID)
	return nil
}

func DeleteUser(id uint) (err error) {
	log.Printf("repository.DeleteUser: Soft deleting user with ID %v\n", id)
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		log.Printf("repository.DeleteUser: Failed to delete user with ID %v, error: %v\n", id, err)
		return err
	}

	log.Printf("repository.DeleteUser: User with ID %v deleted successfully\n", id)
	return nil
}
