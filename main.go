package main

import (
	"Gim/pkg/pmg"
	"fmt"
)

func main() {
	g := pmg.OpenGeorgesHole("./pkg/pmg/test.json")
	fmt.Println(g.Image)
}
