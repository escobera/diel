package resource

import "time"

type Student struct {
	Id           int32     `json:"id"`
	Name         string    `json:"name"`
	Registration string    `json:"registration"`
	CreatedAt    time.Time `json:"created"`
}
