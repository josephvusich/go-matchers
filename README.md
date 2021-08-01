# go-matchers

Ruleset helper for ordered include/exclude logic, with builtin support for the stdlib [flag](https://pkg.go.dev/flag) package. Used by [fdf](https://github.com/josephvusich/fdf).

# Usage

## Parsing with `flag` package

```shell
$ go get github.com/josephvusich/go-matchers
```

```go
package example

import (
	"flag"
	
	"github.com/josephvusich/go-matchers"
	"github.com/josephvusich/go-matchers/glob"
)

func FlagExample() {
	ruleset := matchers.RuleSet{DefaultInclude: true}

	include, exclude := ruleset.FlagValues(glob.NewMatcher)

	fs := flag.NewFlagSet("test", flag.ExitOnError)
	fs.Var(include, "include", "include matching files")
	fs.Var(exclude, "exclude", "exclude matching files")

	fs.Parse([]string{`--exclude`, `*.bar`, `--include`, `foo.bar`})

	ruleset.Includes("foo.bar")  // true
	ruleset.Includes("fizz.bar") // false
}
```

## Custom Matchers

You can extend to patterns beyond simple globs using the provided `matchers.Matcher` interface:

```go
type Matcher interface {
	Match(string) bool
	
	// A string representation of this Matcher.
	// glob.Matcher returns the glob pattern.
	String() string
}
```