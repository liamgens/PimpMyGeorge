package pmg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Georges all my kids
type Georges struct {
	Georges []George `json:"georges"`
}

// George is my pops
type George struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Accessories []Bling
}

//Bling to make him look cute
type Bling struct {
	Noun     string   `json:"noun"`
	Adj      string   `json:"adj"`
	Location Location `json:"location"`
	Image    string   `json:"image"`
}

//Location for the bling
type Location int

//Location enum
const (
	HEAD Location = 0 + iota
	NECK
	TORSO
	LEGS
	HANDS
	FEET
)

// OpenGeorgesHole Pay the troll toll to get in this boys hole
func OpenGeorgesHole(path string) []George {
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var georges Georges

	json.Unmarshal(byteValue, &georges)

	return georges.Georges
}
