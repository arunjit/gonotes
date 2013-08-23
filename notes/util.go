package notes

import (
	"regexp"
	"strconv"
)

type SearchOptions struct {
	All      bool
	ID       int64
	Title    string
	Keywords []string
}

func ParseQuery(query string) *SearchOptions {
	if query == "*" {
		return &SearchOptions{All: true}
	}
	opts := new(SearchOptions)
	parts := parse(query)
	kvre := regexp.MustCompile(`(id|t):(.+)`)
	for _, p := range parts {
		match := kvre.FindStringSubmatch(p)
		if len(match) == 3 {
			switch match[1] {
			case "id":
				if i, err := strconv.Atoi(match[2]); err == nil {
					opts.ID = int64(i)
				}
			case "t":
				opts.Title = match[2]
			}
		} else {
			opts.Keywords = append(opts.Keywords, p)
		}
	}
	return opts
}

func parse(str string) []string {
	var values []string
	var value []rune
	quot := false
	for _, c := range str {
		switch c {
		case '"':
			if !quot {
				quot = true
			} else {
				if string(value) != "" {
					values = append(values, string(value))
				}
				value = []rune{}
				quot = false
			}
		case ' ':
			if !quot {
				if string(value) != "" {
					values = append(values, string(value))
				}
				value = []rune{}
			} else {
				value = append(value, c)
			}
		default:
			value = append(value, c)
		}
	}
	if len(value) > 0 {
		values = append(values, string(value))
	}
	return values
}
