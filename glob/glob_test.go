package glob

import (
	"flag"
	"testing"

	"github.com/josephvusich/go-matchers"
	"github.com/stretchr/testify/require"
)

func TestGlobSetString(t *testing.T) {
	assert := require.New(t)

	gs := &matchers.RuleSet{DefaultInclude: false}

	g, err := NewMatcher("a/**/*")
	assert.NoError(err)
	gs.Add(g, true)
	g, err = NewMatcher("a/**/bar")
	assert.NoError(err)
	gs.Add(g, false)

	assert.Equal("a/**/*\na/**/bar", gs.String())
}

func TestGlobRuleSet(t *testing.T) {
	assert := require.New(t)

	gs := &matchers.RuleSet{DefaultInclude: false}

	g, err := NewMatcher("a/**/*")
	assert.NoError(err)
	gs.Add(g, true)
	g, err = NewMatcher("a/**/bar")
	assert.NoError(err)
	gs.Add(g, false)

	checkCases(assert, gs)
}

func TestGlobFlagValues(t *testing.T) {
	assert := require.New(t)

	fs := flag.NewFlagSet("test", flag.PanicOnError)

	gs := &matchers.RuleSet{DefaultInclude: false}
	incl, excl := gs.FlagValues(NewMatcher)
	fs.Var(incl, "include", "")
	fs.Var(excl, "exclude", "")

	assert.NoError(fs.Parse([]string{
		`--include`, `a/**/*`,
		`--exclude`, `a/**/bar`,
	}))

	checkCases(assert, gs)
}

func checkCases(assert *require.Assertions, gs *matchers.RuleSet) {
	cases := map[string]bool{
		"a/foo":     true,
		"a/foo/bar": false,
		"b/foo":     false,
		"b/foo/bar": false,
	}

	for in, out := range cases {
		assert.Equal(out, gs.Includes(in), "expected Includes=%t for '%s'", out, in)
	}
}
