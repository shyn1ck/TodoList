package repository

import (
	"todoList/db"
	"todoList/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Where("is_deleted = ?", false).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func UpdateUser(user models.User) (err error) {
	err = db.GetDBConn().Save(&user).Error
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) (err error) {
	err = db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Update("is_deleted", true).Error
	if err != nil {
		return err
	}
	return nil
}
