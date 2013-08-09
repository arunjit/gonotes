package notes

import (
	// "fmt"
	"net/http"

	rpc "github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

// The notes server. This includes a JSON-RPC 2.0 server wrapped by a higher
// level, middleware-esque server.

type server struct {
	handler http.Handler
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.handler.ServeHTTP(w, r)
}

// Creates a new server for the NotesService.
func NewServer(svc NotesService) *server {
	s := rpc.NewServer()
	s.RegisterCodec(json2.NewCodec(), "application/json")
	s.RegisterService(svc, "notes")
	return &server{handler: s}
}
