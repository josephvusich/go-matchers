package regex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegex_Match(t *testing.T) {
	assert := require.New(t)

	m1, err := CompileMatcher("xyz")
	assert.NoError(err)

	m2, err := CompileMatcher("^xyz$")
	assert.NoError(err)

	testStr := "abc...xyz"
	assert.True(m1.Match(testStr))
	assert.False(m2.Match(testStr))
}
