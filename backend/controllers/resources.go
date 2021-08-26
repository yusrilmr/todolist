package controllers

import (
	"github.com/yusrilmr/todolist/backend/models"
)

// Models for JSON resources
type (
	// UserResource is model for Post - /user/register
	UserResource struct {
		Data models.User `json:"data"`
	}

	// LoginResource is model for Post - /user/login
	LoginResource struct {
		Data LoginModel `json:"data"`
	}

	// AuthUserResource is response model for authorized user Post - /user/login
	AuthUserResource struct {
		Data AuthUserModel `json:"data"`
	}

	// TaskResource is model for Post/Put - /tasks
	// TaskResource is model for Get - /tasks/id
	TaskResource struct {
		Data models.Task `json:"data"`
	}

	// TasksResource is model for Get - /tasks
	TasksResource struct {
		Data []models.Task `json:"data"`
	}

	// LabelResource is model for Post/Put - /labels
	LabelResource struct {
		Data models.Label `json:"data"`
	}

	// LabelsResource is model for Get - /lables
	// LabelsResource is model for /lables/tasks/id
	LabelsResource struct {
		Data []models.Label `json:"data"`
	}

	// LoginModel is model for authentication
	LoginModel struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// AuthUserModel is model for authorized user with access token
	AuthUserModel struct {
		User  models.User `json:"user"`
		Token string      `json:"token"`
	}
)