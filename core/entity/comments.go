package entity

import (
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	UserId             uint
	CommentDescription string `json:"commentDescription"`
}

type CommentInput struct {
	CommentDescription string `json:"commentDescription"`
}
