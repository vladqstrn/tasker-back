package models

import "time"

type Task struct {
	Id          int
	Title       string
	Text        string
	Description string
	Executor    string
	Status      string
	Created_at  time.Time
}
