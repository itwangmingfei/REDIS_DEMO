package model

type User struct {
	Name string `form:"name" binding:"required"`
} 
