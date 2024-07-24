package main

import (
	"fmt"
	"github.com/CristianC4/toolkit"
)

func main() {
	toSlug := "now is the time 123"
	var tools toolkit.Tools

	slugified, er := tools.Slugify(toSlug)
	if er != nil {
		fmt.Println(er)
	}
	fmt.Println(slugified)
}
