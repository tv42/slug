package slug

import (
	"net/url"
	"strings"
)

// URLString returns a slugified string based on the URL passed as a
// string. This is a convenience wrapper over URL. It fails only if
// parsing the URL fails.
func URLString(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return URL(u), nil
}

const (
	wwwPrefix   = "www."
	indexSuffix = "/index.html"
	htmlSuffix  = ".html"
)

// URL returns a slugified string based on the URL. In addition to the
// usual slugification rules, the following simplifications are done:
//
//   - schemes `http` and `https` are removed
//   - a leading `www.` in hostname is removed
//   - a trailing `/index.html` or `.html` is removed
func URL(u *url.URL) string {
	// take a copy so we don't mutate parent
	tmp := *u
	u = &tmp

	switch u.Scheme {
	case "http", "https":
		u.Scheme = ""
	}

	if strings.HasPrefix(u.Host, wwwPrefix) {
		u.Host = u.Host[len(wwwPrefix):]
	}

	if strings.HasSuffix(u.Path, indexSuffix) {
		u.Path = u.Path[:len(u.Path)-len(indexSuffix)]
	} else if strings.HasSuffix(u.Path, htmlSuffix) {
		u.Path = u.Path[:len(u.Path)-len(htmlSuffix)]
	}

	return Slug(u.String())
}
