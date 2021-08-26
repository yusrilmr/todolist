package data

import (
	"github.com/yusrilmr/todolist/backend/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	C *gorm.DB
}

// CreateUser creates new user
func (r *UserRepository) CreateUser(user *models.User) error {
	hpass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	user.HashPassword = hpass
	// Clear the incoming text password
	user.Password = ""

	err = r.C.Create(&user).Error
	return err
}

// Login validates email and password
func (r *UserRepository) Login(user models.User) (u models.User, err error) {
	err = r.C.Where("email = ?", user.Email).First(&u).Error
	if err != nil {
		return
	}
	// Validate password
	err = bcrypt.CompareHashAndPassword(u.HashPassword, []byte(user.Password))
	if err != nil {
		u = models.User{}
	}
	return
}
