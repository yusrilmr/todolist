package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/yusrilmr/todolist/backend/common"
	"github.com/yusrilmr/todolist/backend/data"
	//"github.com/yusrilmr/todolist/backend/models"
)

// CreateTask insert a new Task
// Handler for HTTP Post - "/tasks
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var dataResource TaskResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid data",
			500,
		)
		return
	}

	task := &dataResource.Data
	context := GetNewContext().PostgresDB
	repo := &data.TaskRepository{C: context}
	repo.Create(task)
	if j, err := json.Marshal(TaskResource{Data: *task}); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(j)
	}
}

// GetTasks returns all Task
// Handler for HTTP Get - "/tasks"
func GetTasks(w http.ResponseWriter, r *http.Request) {
	context := GetNewContext().PostgresDB
	repo := &data.TaskRepository{C: context}
	tasks := repo.GetAll()
	j, err := json.Marshal(TasksResource{Data: tasks})

	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// GetTaskById returns a single Task  by id
// Handler for HTTP Get - "/tasks/{id}"
func GetTaskById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := GetNewContext().PostgresDB
	repo := &data.TaskRepository{C: context}
	task, err := repo.GetById(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	j, err := json.Marshal(task)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// GetTasksByUser returns all Tasks created by a User
// Handler for HTTP Get - "/tasks/users/{id}"
func GetTasksByUser(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	user := vars["id"]

	context := GetNewContext().PostgresDB
	repo := &data.TaskRepository{C: context}

	tasks := repo.GetByUser(user)
	j, err := json.Marshal(TasksResource{Data: tasks})
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// UpdateTask update an existing Task
// Handler for HTTP Put - "/tasks/{id}"
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	var dataResource TaskResource
	// Decode the incoming Task json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Task data",
			500,
		)
		return
	}
	task := &dataResource.Data
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	task.ID = uint(u64)
	context := GetNewContext().PostgresDB
	repo := &data.TaskRepository{C: context}
	// Update an existing Task
	if err := repo.Update(task); err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DeleteTask deelete an existing Task
// Handler for HTTP Delete - "/tasks/{id}"
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := GetNewContext().PostgresDB
	repo := &data.TaskRepository{C: context}
	// Delete an existing Task
	err := repo.Delete(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}