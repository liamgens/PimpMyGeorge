package main

import (
	"fmt"

	"github.com/liamgens/Gim/pkg/pmg"
)

func main() {
	g, b := pmg.FetchGeorgeBlingData("./pkg/pmg/georges.json", "./pkg/pmg/blang.json")
	fmt.Println(b)

	//Iterate over all Georges
	for i := 0; i < len(g); i++ {
		//Iterate over each georges pieces and print it out
		fmt.Println(g[i].Rectals)
	}

}
