package controllers

import (
	"github.com/yusrilmr/todolist/backend/common"
	"gorm.io/gorm"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
	PostgresDB *gorm.DB
}

// Create a new Context object for each HTTP request
func NewContext() *Context {
	db := common.GetDB()
	db_ctx := db.WithContext(db.Statement.Context)
	context := &Context{
		PostgresDB: db_ctx,
	}
	return context
}
