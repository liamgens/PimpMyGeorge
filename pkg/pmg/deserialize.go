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

// Blang for the boys
type Blang struct {
	Blang []Bling `json:"blang"`
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
func OpenGeorgesHole(georgePath string, blingPath string) ([]George, []Bling) {
	georgeJSON, georgeErr := os.Open(georgePath)
	blingJSON, blingErr := os.Open(blingPath)

	if georgeErr != nil {
		fmt.Println(georgeErr)
	}

	if blingErr != nil {
		fmt.Println(blingErr)
	}

	defer georgeJSON.Close()
	defer blingJSON.Close()

	georgeBytes, _ := ioutil.ReadAll(georgeJSON)
	blangBytes, _ := ioutil.ReadAll(blingJSON)

	var georges Georges
	var blang Blang

	json.Unmarshal(georgeBytes, &georges)
	json.Unmarshal(blangBytes, &blang)

	return georges.Georges, blang.Blang
}
