package model

import (
	"time"
)

// Document ...
type Document struct {
	ID               int
	Title            string
	Description      string
	LastModification time.Time
	Content          string
}
