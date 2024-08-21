package repository

import (
	"todoList/db"
	"todoList/logger"
	"todoList/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers]: Failed to fetch users, error: %v\n", err)
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID]: Failed to fetch user, ID: %v, error: %v\n", id, err)
		return user, err
	}
	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername]: Failed to fetch user by username %v, error: %v\n", username, err)
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser]: Failed to create user, error: %v\n", err)
		return err
	}
	return nil
}

func UpdateUser(user models.User) (err error) {
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.UpdateUser]: Failed to update user with ID %v, error: %v\n", user.ID, err)
		return err
	}

	return nil
}

func DeleteUser(id uint) (err error) {
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		logger.Error.Printf("[repository.DeleteUser]: Failed to delete user with ID %v, error: %v\n", id, err)
		return err
	}
	return nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, err
	}
	return user, nil
}
