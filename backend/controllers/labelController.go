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

// CreateLabel insert a new Label document
// Handler for HTTP Post - "/labels
func CreateLabel(w http.ResponseWriter, r *http.Request) {
	var dataResource LabelResource
	// Decode the incoming User json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid User data",
			500,
		)
		return
	}

	label := &dataResource.Data
	context := NewContext().PostgresDB
	repo := &data.LabelRepository{C: context}
	repo.Create(label)
	if j, err := json.Marshal(LabelResource{Data: *label}); err != nil {
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

//// GetLabels returns all Label document
//// Handler for HTTP Get - "/labels"
//func GetLabels(w http.ResponseWriter, r *http.Request) {
//	context := NewContext()
//	defer context.Close()
//	col := context.DbCollection("labels")
//	repo := &data.LabelRepository{C: col}
//	labels := repo.GetAll()
//	j, err := json.Marshal(LabelsResource{Data: labels})
//	if err != nil {
//		common.DisplayAppError(
//			w,
//			err,
//			"An unexpected error has occurred",
//			500,
//		)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(j)
//}
//
// GetLabelById returns a single Label document by id
// Handler for HTTP Get - "/labels/{id}"
func GetLabelById(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext().PostgresDB
	repo := &data.LabelRepository{C: context}
	label, err := repo.GetById(id)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"An unexpected error has occurred",
			500,
		)
		return
	}
	j, err := json.Marshal(label)
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
//
//// GetLabelsByUser returns all Labels created by a User
//// Handler for HTTP Get - "/labels/users/{id}"
//func GetLabelsByUser(w http.ResponseWriter, r *http.Request) {
//	// Get id from the incoming url
//	vars := mux.Vars(r)
//	user := vars["id"]
//	context := NewContext()
//	defer context.Close()
//	col := context.DbCollection("labels")
//	repo := &data.LabelRepository{C: col}
//	labels := repo.GetByUser(user)
//	j, err := json.Marshal(LabelsResource{Data: labels})
//	if err != nil {
//		common.DisplayAppError(
//			w,
//			err,
//			"An unexpected error has occurred",
//			500,
//		)
//		return
//	}
//	w.WriteHeader(http.StatusOK)
//	w.Header().Set("Content-Type", "application/json")
//	w.Write(j)
//}
//
// UpdateLabel update an existing Label document
// Handler for HTTP Put - "/labels/{id}"
func UpdateLabel(w http.ResponseWriter, r *http.Request) {
	// Get id from the incoming url
	vars := mux.Vars(r)
	id := vars["id"]
	var dataResource LabelResource
	// Decode the incoming Label json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Label data",
			500,
		)
		return
	}
	label := &dataResource.Data
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
	label.ID = uint(u64)
	context := NewContext().PostgresDB
	repo := &data.LabelRepository{C: context}
	// Update an existing Label document
	if err := repo.Update(label); err != nil {
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

// DeleteLabel deelete an existing Label document
// Handler for HTTP Delete - "/labels/{id}"
func DeleteLabel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	context := NewContext().PostgresDB
	repo := &data.LabelRepository{C: context}
	// Delete an existing Label
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
