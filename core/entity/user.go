package entity

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Comments []Comment `json:"comments"`
	FullName string    `json:"fullName"`
	Age      int64     `json:"age"`
}

type UserInput struct {
	Comments []CommentInput `json:"comments"`
	FullName string         `json:"fullName"`
	Age      int64          `json:"age"`
}
