package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/yusrilmr/todolist/backend/common"
	"github.com/yusrilmr/todolist/backend/data"
	"github.com/yusrilmr/todolist/backend/models"
)

// Handler for HTTP Post - "/users/register"
// Add a new User document
func Register(w http.ResponseWriter, r *http.Request) {
	var dataResource UserResource
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

	user := &dataResource.Data
	context := NewContext().PostgresDB
	repo := &data.UserRepository{C: context}
	// Insert User document
	repo.CreateUser(user)
	// Clean-up the hashpassword to eliminate it from response
	user.HashPassword = nil
	if j, err := json.Marshal(UserResource{Data: *user}); err != nil {
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

// Handler for HTTP Post - "/users/login"
// Authenticate with username and apssword
func Login(w http.ResponseWriter, r *http.Request) {
	var dataResource LoginResource
	var token string

	// Decode the incoming Login json
	err := json.NewDecoder(r.Body).Decode(&dataResource)
	if err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid Login data",
			500,
		)
		return
	}

	loginModel := dataResource.Data

	loginUser := models.User{
		Email:    loginModel.Email,
		Password: loginModel.Password,
	}

	context := NewContext().PostgresDB
	repo := &data.UserRepository{C: context}

	// Authenticate the login user
	if user, err := repo.Login(loginUser); err != nil {
		common.DisplayAppError(
			w,
			err,
			"Invalid login credentials",
			401,
		)
		return
	} else { //if login is successful
		// Generate JWT token
		token, err = common.GenerateJWT(user.Email)
		if err != nil {
			common.DisplayAppError(
				w,
				err,
				"Error while generating the access token",
				500,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		user.HashPassword = nil
		authUser := AuthUserModel{
			User:  user,
			Token: token,
		}
		j, err := json.Marshal(AuthUserResource{Data: authUser})
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
		w.Write(j)
	}
}
