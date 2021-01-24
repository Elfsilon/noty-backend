package store

import (
	"github.com/Elfsilon/noty-backend/internal/app/model"
)

// UserRep ...
type UserRep struct {
	store *Store
}

// NewUserRep ...
func NewUserRep(store *Store) *UserRep {
	return &UserRep{
		store: store,
	}
}

// Create ...
func (r *UserRep) Create(user *model.User) (*model.User, error) {
	if err := r.store.db.QueryRow(
		"INSERT INTO users (name, photo, email, login, password) VALUES ($1, $2, $3, $4, $5) RETURNING ID",
		user.Name,
		user.Photo,
		user.Email,
		user.Login,
		user.Password,
	).Scan(&user.ID); err != nil {
		return nil, err
	}

	return user, nil
}
