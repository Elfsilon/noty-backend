package store

import (
	"testing"
)

func TestStore(t testing.T, databaseURL string) *Store {
	t.Helper()

	config := NewConfig()
	config.DatabaseURL = databaseURL
	s := New(config)
	if err := s.Open(); err != nil {
		t.Fatal(err)
	}

	return s
}
