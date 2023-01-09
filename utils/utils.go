package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/rohinivsenthil/golang-getting-started/schema"
)

func SaveToDos(todos schema.ToDoList) (returnError error) {
	output, err := json.Marshal(todos)
	if err != nil {
		log.WithError(err).WithField("todos", todos).Error("Failed to convert todos to json")
		errorMsg := fmt.Sprintf("Failed to convert to json: %s", err.Error())
		returnError := errors.New(errorMsg)

		return returnError
	}

	if err = ioutil.WriteFile("data.json", output, 0666); err != nil {
		log.WithError(err).WithField("data", output).Error("Failed to write data to data.json")
		errorMsg := fmt.Sprintf("Failed to write to data.json: %s", err.Error())
		returnError := errors.New(errorMsg)

		return returnError
	}

	return nil
}