package main

import (
	"Gim/pkg/pmg"
	"fmt"
)

func main() {
	g, b := pmg.OpenGeorgesHole("./pkg/pmg/test.json", "./pkg/pmg/test2.json")
	fmt.Println(g, b)
}
