package app

import "time"

type Photo struct {
    ID        uint      `json:"id" gorm:"primary_key"`
    Title     string    `json:"title"`
    Caption   string    `json:"caption"`
    PhotoUrl  string    `json:"photo_url"`
    UserID    uint      `json:"user_id"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}
