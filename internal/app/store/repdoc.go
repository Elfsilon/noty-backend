package store

import (
	"github.com/Elfsilon/noty-backend/internal/app/model"
)

// DocumentRep ...
type DocumentRep struct {
	store *Store
}

// NewDocumentRep ...
func NewDocumentRep(store *Store) *DocumentRep {
	return &DocumentRep{
		store: store,
	}
}

// Create ...
func (r *DocumentRep) Create(user model.Document) (bool, error) {
	return false, nil
}
