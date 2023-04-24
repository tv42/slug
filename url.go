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
//   - percent-encoding is undone
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

	// Take over formatting of Path, Query, Fragment so we can undo any percent escaping for them, for prettier output.
	path := u.EscapedPath()
	u.Path = ""
	// We don't need to care about ForceQuery, it would be deduplicated or trimmed by our dash rules anyway.
	query := u.RawQuery
	u.RawQuery = ""
	fragment := u.EscapedFragment()
	u.Fragment = ""

	var buf strings.Builder
	buf.WriteString(u.String())

	// See `net/url.URL.String` for edge cases.
	if path != "" && path[0] != '/' && u.Host != "" {
		buf.WriteByte('/')
	}
	if buf.Len() == 0 {
		if segment, _, _ := strings.Cut(path, "/"); strings.Contains(segment, ":") {
			buf.WriteString("./")
		}
	}

	// Undo percent-encoding for path
	if p, err := url.PathUnescape(path); err == nil {
		path = p
	}
	buf.WriteString(path)

	// Undo percent-encoding for query
	if q, err := url.QueryUnescape(query); err == nil {
		query = q
	}
	if query != "" {
		buf.WriteByte('?')
		buf.WriteString(query)
	}

	// Undo percent-encoding for fragment.
	// The difference between `QueryUnescape` and `PathUnescape` makes no difference here, since both will become dashes.
	if f, err := url.QueryUnescape(fragment); err == nil {
		fragment = f
	}
	if fragment != "" {
		buf.WriteByte('#')
		buf.WriteString(fragment)
	}

	return Slug(buf.String())
}
