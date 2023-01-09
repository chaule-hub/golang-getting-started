package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rohinivsenthil/golang-getting-started/schema"
	"github.com/rohinivsenthil/golang-getting-started/utils"

	log "github.com/sirupsen/logrus"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	log.Info("Received get requests")

	// read from file
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		// return error in response
		log.WithError(err).Error("Failed to read from data.json")
		fmt.Fprintf(w, "Failed to get data: %s", err.Error())
		return
	}

	// set response
	w.Write(data)
}

func CreateNewTodo(w http.ResponseWriter, r *http.Request) {
	log.Info("Received post requests")

	// read POST request body
	var data schema.ToDoDO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Failed to read data from post")
		fmt.Fprintf(w, "Failed to get data: %s", err.Error())
		return
	}

	// read existing data from file
	existing_data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.WithError(err).Error("Failed to read from data.json")
		fmt.Fprintf(w, "Failed to get existing data: %s", err.Error())
	}

	// parse existing string data to json
	var todos schema.ToDoList
	if err = json.Unmarshal(existing_data, &todos); err != nil {
		log.WithError(err).Error("Failed to parse existing data")
		fmt.Fprintf(w, "Failed to parse existing data: %s", err.Error())
		return
	}

	// add new data
	todos = append(todos, schema.ToDo{Id: uuid.NewString(), Text: data.Text})
	if err := utils.SaveToDos(todos); err != nil {
		fmt.Fprintf(w, "Failed to save data: %s", err.Error())
		return
	}

	// set response
	w.Write([]byte("Successfully wrote data"))
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	log.Info("Received PUT requests")

	vars := mux.Vars(r)
	updateId := vars["todoId"]

	// read PUT request body
	var data schema.ToDoDO
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		log.WithError(err).Error("Failed to read data from PUT")
		fmt.Fprintf(w, "Failed to get data: %s", err.Error())
		return
	}

	// read existing data from file
	existing_data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.WithError(err).Error("Failed to read from data.json")
		fmt.Fprintf(w, "Failed to get existing data: %s", err.Error())
	}

	// parse existing string data to json
	var todos schema.ToDoList
	if err = json.Unmarshal(existing_data, &todos); err != nil {
		log.WithError(err).Error("Failed to parse existing data")
		fmt.Fprintf(w, "Failed to parse existing data: %s", err.Error())
		return
	}

	var i int
	for i = 0; i < len(todos); i++ {
		todo := todos[i]
		if todo.Id == updateId {
			todos[i].Text = data.Text
			break
		}
	}

	if i == len(todos) {
		fmt.Fprintf(w, "Todo with ID %s not found", updateId)
		return
	}

	// write to file
	if err := utils.SaveToDos(todos); err != nil {
		fmt.Fprintf(w, "Failed to save data: %s", err.Error())
		return
	}

	// set response
	w.Write([]byte("Successfully updated data"))
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	log.Info("Received PUT requests")

	vars := mux.Vars(r)
	deleteId := vars["todoId"]

	// read existing data from file
	existing_data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.WithError(err).Error("Failed to read from data.json")
		fmt.Fprintf(w, "Failed to get existing data: %s", err.Error())
	}

	// parse existing string data to json
	var todos schema.ToDoList
	if err = json.Unmarshal(existing_data, &todos); err != nil {
		log.WithError(err).Error("Failed to parse existing data")
		fmt.Fprintf(w, "Failed to parse existing data: %s", err.Error())
		return
	}

	var i int
	for i = 0; i < len(todos); i++ {
		todo := todos[i]
		if todo.Id == deleteId {
			break
		}
	}

	if i == len(todos) {
		fmt.Fprintf(w, "Todo with ID %s not found", deleteId)
		return
	}

	todos = append(todos[:i], todos[i+1:]...)
	if err := utils.SaveToDos(todos); err != nil {
		fmt.Fprintf(w, "Failed to save data: %s", err.Error())
		return
	}

	// set response
	w.Write([]byte("Successfully deleted data"))
}