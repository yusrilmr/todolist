package controllers

import (
	"github.com/yusrilmr/todolist/backend/common"
	"gorm.io/gorm"
)

// Struct used for maintaining HTTP Request Context
type Context struct {
	PostgresDB *gorm.DB
}

// GetNewContext creates a new db context object
func GetNewContext() *Context {
	db := common.GetDB()
	dbCtx := db.WithContext(db.Statement.Context)
	context := &Context{
		PostgresDB: dbCtx,
	}
	return context
}
