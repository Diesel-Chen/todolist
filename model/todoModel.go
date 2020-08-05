package model

import "github.com/jinzhu/gorm"

type (
	// TodoModel里对应数据库的字段
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
	// TransformedTodo 是指定数据返回前端
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)
