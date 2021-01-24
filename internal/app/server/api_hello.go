package server

import (
	"io"
	"net/http"
)

func (s *Server) helloFunc() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello")
	}
}
