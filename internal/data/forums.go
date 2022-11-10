// Filename: internal/data/forums.go

package data

import (
	"time"
)

type Forum struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Message   string    `json:"message"`
	User      string    `json:"user"`
	Version   int32     `json:"version"`
}