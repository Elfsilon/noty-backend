package store

import (
	"database/sql"

	//Postgres driver
	_ "github.com/lib/pq"
)

// Store ...
type Store struct {
	config *Config
	db     *sql.DB

	// User model interface
	User *UserRep
}

func (s *Store) initRepos() {
	s.User = NewUserRep(s)
}

// New ...
func New(config *Config) *Store {
	store := Store{
		config: config,
	}
	store.initRepos()
	return &store
}

// Open ...
func (s *Store) Open() error {
	conn, err := sql.Open("postgres", s.config.DatabaseURL)
	if err != nil {
		return err
	}

	if err := conn.Ping(); err != nil {
		return err
	}

	s.db = conn

	return nil
}

// Close ...
func (s *Store) Close() {
	s.db.Close()
}
