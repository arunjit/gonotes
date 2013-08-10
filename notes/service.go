package notes

import (
	"net/http"
)

type NotesDatabase interface {
	Search(*http.Request, *SearchOptions) ([]*Note, error)
	Update(*http.Request, *Note) (*Note, error)
	Delete(*http.Request, int64) error
}

type NotesService struct {
	db NotesDatabase
}

func NewNotesService(db NotesDatabase) *NotesService {
	return &NotesService{db}
}

func (s *NotesService) Search(r *http.Request, args *SearchArgs, result *SearchResult) error {
	opts := ParseQuery(args.Query)
	// opts.Offset = args.Offset
	// opts.Limit = args.Limit
	notes, err := s.db.Search(r, opts)
	if err != nil {
		return err
	}
	result.Notes = notes
	return nil
}

func (s *NotesService) Update(r *http.Request, args *Note, result *Note) error {
	note, err := s.db.Update(r, args)
	if err != nil {
		return err
	}
	result.ID = note.ID
	result.Title = note.Title
	return nil
}

func (s *NotesService) Delete(r *http.Request, args *DeleteArgs, result *NoContent) error {
	if err := s.db.Delete(r, args.ID); err != nil {
		return err
	}
	return nil
}
