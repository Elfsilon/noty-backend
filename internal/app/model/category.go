package model

import (
	"time"
)

// Category ...
type Category struct {
	ID               int
	Title            string
	Description      string
	LastModification time.Time
}
