package notes

import (
	// "fmt"
	"net/http"

	rpc "github.com/gorilla/rpc/v2"
	"github.com/gorilla/rpc/v2/json2"
)

const (
	RPCPath = "/rpc"
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
func NewServer(svc *NotesService) *server {
	svr := rpc.NewServer()
	svr.RegisterCodec(json2.NewCodec(), "application/json")
	svr.RegisterService(svc, "notes")
	return &server{handler: svr}
}
