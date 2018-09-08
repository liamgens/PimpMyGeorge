package main

import (
	"fmt"

	"github.com/liamgens/Gim/pkg/pmg"
)

func main() {
	g, b := pmg.FetchGeorgeBlingData("./pkg/pmg/test.json", "./pkg/pmg/test2.json")
	fmt.Println(g, b)
}
