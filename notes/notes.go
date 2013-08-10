package notes

// The Note type
type Note struct {
	ID    int64  `json:"id" datastore:"-"`
	Title string `json:"title,omitempty" datastore:"t"`
}

// `Search` is used for all data retreival.
// The `Query` parameter can parse the following tokens:
//     *            =>  list all
//     id:1234      =>  list #1234
//     t:xyz        =>  list all with Title ^~ xyz
//     t:"xyz abc"  =>  list all with Title == "xyz abc"
type SearchArgs struct {
	Query  string `json:"query"`
	Offset int32  `json:"offset,omitempty"`
	Limit  int32  `json:"limit,omitempty"`
}

type SearchResult struct {
	Notes []*Note `json:"notes"`
}

// Create/update requests directly use `Note`.

type DeleteArgs struct {
	ID int64 `json:"id"`
}

type NoContent struct{}
