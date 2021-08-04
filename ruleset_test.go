package matchers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRuleSetString(t *testing.T) {
	assert := require.New(t)

	ruleset := &RuleSet{}

	assert.Equal("default: exclude", ruleset.String())

	ruleset.Add(&mockMatcher{s: "foo"}, true)
	ruleset.Add(&mockMatcher{s: "bar"}, false)

	assert.Equal("default: exclude\ninclude: foo\nexclude: bar", ruleset.String())

	ruleset.DefaultInclude = true

	assert.Equal("default: include\ninclude: foo\nexclude: bar", ruleset.String())
}

type mockMatcher struct {
	s string
}

func (m *mockMatcher) Match(string) bool {
	return false
}

func (m *mockMatcher) String() string {
	return m.s
}
