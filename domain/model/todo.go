package model

import "time"

type TodoId int64
type Title string
type Description string

type Todo struct {
	Id          TodoId      `json:"id"`
	Title       Title       `json:"title"`
	Description Description `json:"description"`
	CreatedAt   time.Time   `json:"created_at"`
	DeleteAt    time.Time   `json:"delete_at"`
}
