package main

import (
	"fmt"

	"github.com/liamgens/PimpMyGeorge/pmg"
)

func main() {
	g, b := pmg.FetchGeorgeBlingData("./pmg/georges.json", "./pmg/blang.json")
	fmt.Println(b)

	//Iterate over all Georges
	for i := 0; i < len(g); i++ {
		//Iterate over each georges pieces and print it out
		fmt.Println(g[i].Rectals)
	}

}
