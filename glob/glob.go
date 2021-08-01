package glob

import (
	"github.com/josephvusich/go-matchers"
	"github.com/mattn/go-zglob"
)

type glob struct {
	pattern string
}

func NewMatcher(pattern string) (r matchers.Matcher, err error) {
	return &glob{
		pattern: pattern,
	}, nil
}

func (g *glob) Match(path string) bool {
	ok, err := zglob.Match(g.pattern, path)
	if err != nil {
		panic(err)
	}
	return ok
}

func (g *glob) String() string {
	return g.pattern
}
