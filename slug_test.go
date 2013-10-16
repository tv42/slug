package slug_test

import (
	"github.com/tv42/slug"
	"testing"
)

func TestSlug(t *testing.T) {
	tests := []struct {
		In, Want string
	}{
		{"foo bar", "foo-bar"},
		{"foo  bar", "foo-bar"},
		{"foo   ", "foo"},
		{"exam҉ple", "example"},
		{"Foo Bar", "foo-bar"},
		{".foo", "foo"},
		{"../evil", "evil"},
		{"../../etc/passwd", "etc-passwd"},
	}

	for _, test := range tests {
		got := slug.Slug(test.In)
		if got != test.Want {
			t.Errorf("wrong slug: URL(%q)=%q want %q", test.In, got, test.Want)
		}
	}
}
