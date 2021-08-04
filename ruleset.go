package matchers

import (
	"flag"
	"fmt"
	"strings"
)

type RuleSet struct {
	DefaultInclude bool
	rules          []*rule
}

type rule struct {
	Matcher
	include bool
}

type Matcher interface {
	Match(string) bool
	String() string
}

type MatcherFunc func(input string) (Matcher, error)

func (s *RuleSet) Add(r Matcher, include bool) {
	s.rules = append(s.rules, &rule{Matcher: r, include: include})
}

type flagValue struct {
	gs      *RuleSet
	include bool
	factory MatcherFunc
}

func (f *flagValue) Set(input string) error {
	r, err := f.factory(input)
	if err != nil {
		return err
	}
	f.gs.Add(r, f.include)
	return nil
}

func (f *flagValue) String() string {
	if f.gs == nil {
		return ""
	}
	return f.gs.String()
}

// For integration with flag.Var
func (s *RuleSet) FlagValues(f MatcherFunc) (include, exclude flag.Value) {
	return &flagValue{
			gs:      s,
			include: true,
			factory: f,
		}, &flagValue{
			gs:      s,
			include: false,
			factory: f,
		}
}

func ruleString(b bool) string {
	if b {
		return "include"
	}
	return "exclude"
}

func (s *RuleSet) String() string {
	elems := make([]string, 1, len(s.rules)+1)
	elems[0] = fmt.Sprintf("default: %s", ruleString(s.DefaultInclude))
	for _, r := range s.rules {
		elems = append(elems, fmt.Sprintf("%s: %s", ruleString(r.include), r.Matcher.String()))
	}
	return strings.Join(elems, "\n")
}

// A later Matcher takes precedence over an earlier one relative to the order added.
// Default (before matching any rules) is the opposite of the first Include/Exclude type.
// E.g., a list beginning with Include has an implicit "Exclude all" base rule, and vice versa.
// Empty RuleSet returns DefaultInclude.
func (s *RuleSet) Includes(check string) bool {
	if len(s.rules) == 0 {
		return s.DefaultInclude
	}

	include := !s.rules[0].include
	for _, r := range s.rules {
		if r.Match(check) {
			include = r.include
		}
	}
	return include
}
