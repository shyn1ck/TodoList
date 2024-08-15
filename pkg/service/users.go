package service

import (
	"errors"
	"gorm.io/gorm"
	"todoList/models"
	"todoList/pkg/repository"
	"todoList/utils"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserById(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user models.User) error {
	_, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	user.Password = utils.GenerateHash(user.Password)
	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUser(user models.User) error {
	err := repository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	err := repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
