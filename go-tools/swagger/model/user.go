package model

import "time"

// User 用户模型
type User struct {
	ID        uint      `json:"id" example:"1"`
	Name      string    `json:"name" example:"张三"`
	Email     string    `json:"email" example:"zhangsan@example.com"`
	Age       int       `json:"age" example:"25"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateUserRequest 创建用户请求
type CreateUserRequest struct {
	Name  string `json:"name" binding:"required" example:"张三"`
	Email string `json:"email" binding:"required,email" example:"zhangsan@example.com"`
	Age   int    `json:"age" binding:"required,min=1,max=150" example:"25"`
}

// UpdateUserRequest 更新用户请求
type UpdateUserRequest struct {
	Name  string `json:"name" example:"李四"`
	Email string `json:"email" binding:"email" example:"lisi@example.com"`
	Age   int    `json:"age" binding:"min=1,max=150" example:"30"`
}

// UserListResponse 用户列表响应
type UserListResponse struct {
	Users []User `json:"users"`
	Total int    `json:"total" example:"100"`
}
