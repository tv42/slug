package slug_test

import (
	"github.com/tv42/slug"
	"net/url"
	"testing"
)

func TestURLString(t *testing.T) {
	tests := []struct {
		In, Want string
	}{
		{"http://example.com/", "example-com"},
		{"http://example.com/foo/bar", "example-com-foo-bar"},
		{"http://www.example.com/", "example-com"},
		{"https://www.example.com/", "example-com"},
		{"http://ex...am҉ple.com/", "ex-am-d2-89ple-com"},
		{"/foo", "foo"},
		{"//foo", "foo"},
		{"http:foo", "foo"},
		{"isbn:foobar", "isbn-foobar"},
		{"http:///foo", "foo"},
		{"http://www.example.com/~u/foo", "example-com-u-foo"},
		{"http://www.example.com/?foo=bar&baz=thud", "example-com-foo-bar-baz-thud"},
		{"http://jdoe:sekrit@www.example.com/?foo=bar&baz=thud", "jdoe-sekrit-example-com-foo-bar-baz-thud"},
		{"http://www.example.com/foo#bar", "example-com-foo-bar"},
		{"http://www.example.com/foo/index.html", "example-com-foo"},
		{"http://www.example.com/foo/bar.html", "example-com-foo-bar"},
	}

	for _, test := range tests {
		got, err := slug.URLString(test.In)
		if err != nil {
			t.Errorf("slug URL parse error: URLString(%q) gave %v", test.In, err)
			continue
		}
		if got != test.Want {
			t.Errorf("wrong slug: URLString(%q)=%q want %q", test.In, got, test.Want)
		}
	}
}

func TestURLNoMutate(t *testing.T) {
	in := "http://example.com/foo"
	u, err := url.Parse(in)
	if err != nil {
		t.Fatalf("test internal error: %v", err)
	}
	want := "example-com-foo"
	got := slug.URL(u)
	if got != want {
		t.Errorf("wrong slug: URL(%q)=%q want %q", in, got, want)
	}

	us := u.String()
	if us != in {
		t.Errorf("url got mutated: %q != %q", us, in)
	}
}
