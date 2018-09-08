package PimpMyGeorge;

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Bling struct {
	adjective string
	noun string
}

type George struct {
	// Tell me more tell me more
	bling []Bling
}

func loadGeorge() George {
	return George{}
}

func PimpHim() {

	fmt.Println("Pimping George")

	reader := bufio.NewReader(os.Stdin)
	george := loadGeorge()

	for {
		fmt.Print("$ ")

		input, _ := reader.ReadString('\n')
		input_spl := strings.Split(input, " ")
		george.bling = append(george.bling,
			Bling{input_spl[0], input_spl[1]})

		fmt.Println("George currently is wearing:")
		for i,el := range george.bling {
			fmt.Printf("%d A %s %s", i, el.adjective, el.noun)
		}
	}
}
