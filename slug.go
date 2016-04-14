package slug

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

// don't even quote these
var skipRanges = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var safeRanges = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

func safe(r rune) rune {
	switch {
	case unicode.IsOneOf(skipRanges, r):
		return -1
	case unicode.IsOneOf(safeRanges, r):
		return unicode.ToLower(r)
	}
	return '-'
}

var doubleDashRE = regexp.MustCompile("--+")

func noRepeat(s string) string {
	return doubleDashRE.ReplaceAllString(s, "-")
}

// Slug returns a slugified string. The result will only contain
// lowercase letters, digits and dashes. It will not begin or end with
// a dash, and it will not contain runs of multiple dashes.
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
