package store

import (
	"testing"

	"github.com/Elfsilon/noty-backend/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUserRep_Create(t *testing.T) {
	s := TestStore(t, "databaseURL")

	u, err := s.User.Create(&model.User{
		Name:     "Maks",
		Email:    "Maxucks@gmail.com",
		Login:    "Maxucks",
		Password: "SomePassw0rd",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
