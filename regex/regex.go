package regex

import (
	"errors"
	"regexp"

	"github.com/josephvusich/go-matchers"
)

type regex struct {
	rgx *regexp.Regexp
}

func NewMatcher(r *regexp.Regexp) (m matchers.Matcher, err error) {
	if r == nil {
		return nil, errors.New("regexp must be non-nil")
	}

	return &regex{
		rgx: r,
	}, nil
}

func CompileMatcher(pattern string) (m matchers.Matcher, err error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return NewMatcher(r)
}

func (r *regex) Match(s string) bool {
	return r.rgx.MatchString(s)
}

func (r *regex) String() string {
	return r.rgx.String()
}
