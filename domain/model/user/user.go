package model

import "time"

type UserId int64

type User struct {
	Id        UserId
	Name      string
	CreatedAt time.Time
}
