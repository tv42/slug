package slug

import (
	"code.google.com/p/go.text/unicode/norm"
	"regexp"
	"strings"
	"unicode"
)

// don't even quote these
var _SKIP = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var _SAFE = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

func safe(r rune) rune {
	switch {
	case unicode.IsOneOf(_SKIP, r):
		return -1
	case unicode.IsOneOf(_SAFE, r):
		return unicode.ToLower(r)
	}
	return '-'
}

var _DOUBLEDASH_RE = regexp.MustCompile("--+")

func noRepeat(s string) string {
	return _DOUBLEDASH_RE.ReplaceAllString(s, "-")
}

// Slugify a string. The result will only contain lowercase letters,
// digits and dashes. It will not begin or end with a dash, and it
// will not contain runs of multiple dashes.
//
// It is NOT forced into being ASCII, but may contain any Unicode
// characters, with the above restrictions.
func Slug(s string) string {
	s = norm.NFKD.String(s)
	s = strings.Map(safe, s)
	s = strings.Trim(s, "-")
	s = noRepeat(s)
	return s
}
