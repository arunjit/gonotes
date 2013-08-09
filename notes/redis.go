// +build !appengine

package notes

import (
	"math/rand"
	"net/http"
	"time"
)

type RedisNotesService struct{}

func (s *RedisNotesService) Search(r *http.Request, args *SearchArgs, result *SearchResult) error {
	result.Notes = []Note{}
	return nil
}

func (s *RedisNotesService) Update(r *http.Request, args *Note, result *Note) error {
	id := args.ID
	if id == 0 {
		id = createID()
	}
	// set(id, encode(args))
	result.ID = id
	result.Title = args.Title
	return nil
}

func (s *RedisNotesService) Delete(r *http.Request, args *DeleteArgs, result *DeleteResult) error {
	return nil
}

func createID() int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63()
}
