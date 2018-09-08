package pmg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Head
// Neck
// Torso
// Legs
// Hands
// Feet

/*

George
{
	name: "forman"
	image: base64

	some sort of x,y for placing images
	head: (x1,y1), (x2,y2)

}

Swag
{
	name: "red hat" | "big ass chain"
	location: HEAD | NECK | ... | FEET
	image: base64
}
*/

type George struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// type Bling struct {
// 	noun     string   `json:"noun"`
// 	adj      string   `json:"adj"`
// 	location Location `json:"location"`
// 	image    string   `json:"image"`
// }

type Location int

const (
	HEAD Location = 0 + iota
	NECK
	TORSO
	LEGS
	HANDS
	FEET
)

// OpenGeorgesHole Pay the troll toll to get in this boys hole
func OpenGeorgesHole(path string) George {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var george George

	json.Unmarshal(byteValue, &george)

	return george

}
