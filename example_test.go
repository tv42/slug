package slug_test

import (
	"fmt"
	"github.com/tv42/slug"
)

func ExampleSlug() {
	fmt.Println(slug.Slug("Rødgrød med fløde"))
	fmt.Println(slug.Slug("BUY NOW!!!11eleven"))
	fmt.Println(slug.Slug("../../etc/passwd"))
	// Output:
	// rødgrød-med-fløde
	// buy-now-11eleven
	// etc-passwd
}

func ExampleURLString() {
	s, err := slug.URLString("https://www.example.com/foo/index.html")
	fmt.Println(s, err)
	// Output:
	// example-com-foo <nil>
}
