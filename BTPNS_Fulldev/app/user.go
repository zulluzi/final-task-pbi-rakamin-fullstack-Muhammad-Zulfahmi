package app

import "time"

type User struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    Username  string    `json:"username" binding:"required"`
    Email     string    `json:"email" binding:"required,email"`
    Password  string    `json:"password" binding:"required,min=6"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
