// +build appengine

package notes

import (
	"net/http"
)

// All App Engine specific handling. Uses `appengine/datastore` to store data.
// Initializes the RPC service and the HTTP handler server.

type Datastore struct{}

const (
	kind = "N"
)

func (s *Datastore) Search(r *http.Request, opts *SearchOptions) ([]*Note, error) {
	return []*Note{}, nil
}

func (s *Datastore) Update(r *http.Request, note *Note) (*Note, error) {
	return note, nil
}

func (s *Datastore) Delete(r *http.Request, id int64) error {
	return nil
}

func init() {
	db := new(Datastore)
	svc := NewNotesService(db)
	server := NewServer(svc)
	http.Handle(RPCPath, server)
}
