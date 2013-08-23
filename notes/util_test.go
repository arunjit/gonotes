package notes

import (
	"testing"
)

// func compare(t *testing.T, expected, actual SearchOptions) {
// 	if expected != actual {
// 		t.Errorf("Expected <<%v>>, got <<%v>>\n", expected, actual)
// 	}
// }

func check(t *testing.T, opts *SearchOptions, cond bool) {
	if !cond {
		t.Errorf("Invalid options: %v\n", opts)
	}
}

func TestParseQuery_ListAll(t *testing.T) {
	opts := ParseQuery("*")
	check(t, opts, opts.All == true)
	check(t, opts, opts.ID == 0)
	check(t, opts, opts.Title == "")
	check(t, opts, len(opts.Keywords) == 0)
}

func TestParseQuery_ID(t *testing.T) {
	opts := ParseQuery("id:293")
	check(t, opts, opts.All == false)
	check(t, opts, opts.ID == 293)
	check(t, opts, opts.Title == "")
	check(t, opts, len(opts.Keywords) == 0)
}

func TestParseQuery_Title(t *testing.T) {
	opts := ParseQuery("t:Something")
	check(t, opts, opts.All == false)
	check(t, opts, opts.ID == 0)
	check(t, opts, opts.Title == "Something")
	check(t, opts, len(opts.Keywords) == 0)
}

func TestParseQuery_TitleWithSpaces(t *testing.T) {
	opts := ParseQuery("t:\"Something more\"")
	check(t, opts, opts.All == false)
	check(t, opts, opts.ID == 0)
	check(t, opts, opts.Title == "Something more")
	check(t, opts, len(opts.Keywords) == 0)
}

func TestParseQuery_Mixed(t *testing.T) {
	opts := ParseQuery("id:293 t:\"Something, something\" more words")
	check(t, opts, opts.All == false)
	check(t, opts, opts.ID == 293)
	check(t, opts, opts.Title == "Something, something")
	check(t, opts, opts.Keywords[0] == "more")
	check(t, opts, opts.Keywords[1] == "words")
}
