package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"

	"github.com/tv42/slug"
)

var prog = filepath.Base(os.Args[0])

func usage() {
	fmt.Fprintf(os.Stderr, "Usage of %s:\n", prog)
	fmt.Fprintf(os.Stderr, "  %s STRING..\n", prog)
	fmt.Fprintf(os.Stderr, "\n")
	flag.PrintDefaults()
}

func main() {
	log.SetFlags(0)
	log.SetPrefix(prog + ": ")
	noURL := flag.Bool("no-url", false, "Disable special handling for URLs")
	flag.Usage = usage
	flag.Parse()

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(2)
	}
	for _, arg := range flag.Args() {
		if !*noURL {
			if u, err := url.Parse(arg); err == nil {
				fmt.Println(slug.URL(u))
				continue
			}
		}
		fmt.Println(slug.Slug(arg))
	}
}
