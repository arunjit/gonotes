package notes

type SearchOptions struct{}

func ParseQuery(query string) *SearchOptions {
	opts := new(SearchOptions)
	return opts
}
