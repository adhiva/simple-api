package models

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Task "Object
type Todo struct {
	ID        uuid.UUID `json:"id"`
	Title     string    `json:"title" binding:"required"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (task *Todo) BeforeCreate(scope *gorm.DB) error {
	scope.Statement.SetColumn("CreatedAt", time.Now())
	scope.Statement.SetColumn("ID", uuid.NewV4().String())
	return nil
}

func (task *Todo) BeforeUpdate(scope *gorm.DB) error {
	scope.Statement.SetColumn("UpdatedAt", time.Now())
	return nil
}
