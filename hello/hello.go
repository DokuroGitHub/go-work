package main

import (
	"fmt"

	"github.com/DokuroGitHub/go-crawl"
	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))

	crawl.Crawl(100, 5, 1, 5)
}
