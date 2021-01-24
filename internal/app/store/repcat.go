package store

import (
	"github.com/Elfsilon/noty-backend/internal/app/model"
)

// CategoryRep ...
type CategoryRep struct {
	store *Store
}

// NewCategoryRep ...
func NewCategoryRep(store *Store) *CategoryRep {
	return &CategoryRep{
		store: store,
	}
}

// Create ...
func (r *CategoryRep) Create(user model.Category) (bool, error) {
	return false, nil
}
