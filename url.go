package slug

import (
	"net/url"
	"strings"
)

// Slugify a URL passed as a string. This is a convenience wrapper
// over URL. It fails only if parsing the URL fails.
func URLString(s string) (string, error) {
	u, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	return URL(u), nil
}

const _WWW_PREFIX = "www."
const _INDEX_SUFFIX = "/index.html"
const _HTML_SUFFIX = ".html"

// Slugify a URL. In addition to the usual slugification rules, the
// following simplifications are done:
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

	if strings.HasPrefix(u.Host, _WWW_PREFIX) {
		u.Host = u.Host[len(_WWW_PREFIX):]
	}

	if strings.HasSuffix(u.Path, _INDEX_SUFFIX) {
		u.Path = u.Path[:len(u.Path)-len(_INDEX_SUFFIX)]
	} else if strings.HasSuffix(u.Path, _HTML_SUFFIX) {
		u.Path = u.Path[:len(u.Path)-len(_HTML_SUFFIX)]
	}

	return Slug(u.String())
}
