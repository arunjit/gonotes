// +build appengine

package notes

import (
	"net/http"
)

type DatastoreNotesService struct{}

func (s *DatastoreNotesService) Search(r *http.Request, args *SearchArgs, result *SearchResult) error {
	return nil
}

func (s *DatastoreNotesService) Update(r *http.Request, args *Note, result *Note) error {
	return nil
}

func (s *DatastoreNotesService) Delete(r *http.Request, args *DeleteArgs, result *DeleteResult) error {
	return nil
}

func init() {
	server := NewServer(new(DatastoreNotesService))
	http.Handle("/rpc", server)
}
