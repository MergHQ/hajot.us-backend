package domain

import "time"

type Post struct {
	Id uint
	Content string
	Timestamp time.Time
}
