package notes

import (
	"net/http"
)

type NotesService interface {
	Search(*http.Request, *SearchArgs, *SearchResult) error
	Update(*http.Request, *Note, *Note) error
	Delete(*http.Request, *DeleteArgs, *DeleteResult) error
}
