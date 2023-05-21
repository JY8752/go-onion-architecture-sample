package model

import "time"

type UserId int64

type User struct {
	Id        UserId    `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}
