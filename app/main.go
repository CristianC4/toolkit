package main

import (
	"fmt"
	"github.com/CristianC4/toolkit"
)

func main() {
	var tools toolkit.Tools
	s := tools.RamdonsString(10)
	// import toolkit doesn't add lib on go.mod because we use workspaces
	fmt.Println(s)
}
