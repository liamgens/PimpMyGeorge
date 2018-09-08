package main

import (
	"fmt"

	"github.com/liamgens/Gim/pkg/pmg"
)

func main() {
	g, b := pmg.FetchGeorgeBlingData("./pkg/pmg/georges.json", "./pkg/pmg/blang.json")
	fmt.Println(g, b)
}
