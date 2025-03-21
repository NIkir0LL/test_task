package model

type User struct {
    ID        int    `json:"id"`
    Name      string `json:"name" binding:"required"`
    Email     string `json:"email" binding:"required,email"`
    CreatedAt string `json:"created_at"`
}