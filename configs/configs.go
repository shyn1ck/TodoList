package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"todoList/models"
)

var AppSettings models.Configs

func ReadSettings() error {
	fmt.Println("Starting reading settings file")
	configFile, err := os.Open("configs/configs.json")
	if err != nil {
		return errors.New(fmt.Sprintf("Couldn't open config file. Error is: %s", err.Error()))
	}

	defer func(configFile *os.File) {
		err = configFile.Close()
		if err != nil {
			log.Fatal("Couldn't close config file. Error is: ", err.Error())
		}
	}(configFile)

	fmt.Println("Starting decoding settings file")
	if err = json.NewDecoder(configFile).Decode(&AppSettings); err != nil {
		return errors.New(fmt.Sprintf("Couldn't decode settings json file. Error is: %s", err.Error()))
	}

	return nil
}
