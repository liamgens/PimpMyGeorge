package pmg

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"os"
)

// Georges struct to hold all different Georges
type Georges struct {
	Georges []George `json:"georges"`
}

// Blang struct to hold the bling for the Georges to wear
type Blang struct {
	Blang []Bling `json:"blang"`
}

// George struct for George and all his sick accessories
type George struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	Accessories []Bling
	Pieces      GeorgePieces `json:"pieces"`
	Rectals     map[string]image.Rectangle
}

// GeorgePieces which are real cute and adorable
type GeorgePieces struct {
	Head  [4]int `json:"head"`
	Face  [4]int `json:"face"`
	Neck  [4]int `json:"neck"`
	Torso [4]int `json:"torso"`
	Legs  [4]int `json:"legs"`
	Hands [4]int `json:"hands"`
	Feet  [4]int `json:"feet"`
}

//Bling struct for bling that George can wear
type Bling struct {
	Noun     string `json:"noun"`
	Adj      string `json:"adj"`
	Location string `json:"location"`
	Image    string `json:"image"`
}

// Location enum for where the bling should be placed on George
const (
	HEAD  string = "HEAD"
	FACE  string = "FACE"
	NECK  string = "NECK"
	TORSO string = "TORSO"
	LEGS  string = "LEGS"
	HANDS string = "HANDS"
	FEET  string = "FEET"
)

// FetchGeorgeBlingData Parses JSON files and returns slices of the George structs and Bling structs
func FetchGeorgeBlingData(georgesPath string, blangPath string) ([]George, []Bling) {
	georgesJSON, georgesErr := os.Open(georgesPath)
	blangJSON, blangErr := os.Open(blangPath)

	if georgesErr != nil {
		fmt.Println(georgesErr)
	}

	if blangErr != nil {
		fmt.Println(blangErr)
	}

	defer georgesJSON.Close()
	defer blangJSON.Close()

	georgesBytes, _ := ioutil.ReadAll(georgesJSON)
	blangBytes, _ := ioutil.ReadAll(blangJSON)

	var georges Georges
	var blang Blang

	json.Unmarshal(georgesBytes, &georges)
	json.Unmarshal(blangBytes, &blang)

	for i := 0; i < len(georges.Georges); i++ {
		georges.Georges[i] = populateMap(georges.Georges[i])
	}

	return georges.Georges, blang.Blang
}

func listToRect(coord [4]int) image.Rectangle {
	return image.Rect(coord[0], coord[1], coord[2], coord[3])
}

func populateMap(g George) George {
	g.Rectals = make(map[string]image.Rectangle)
	g.Rectals[HEAD] = listToRect(g.Pieces.Head)
	g.Rectals[FACE] = listToRect(g.Pieces.Face)
	g.Rectals[NECK] = listToRect(g.Pieces.Neck)
	g.Rectals[TORSO] = listToRect(g.Pieces.Torso)
	g.Rectals[LEGS] = listToRect(g.Pieces.Legs)
	g.Rectals[HANDS] = listToRect(g.Pieces.Hands)
	g.Rectals[FEET] = listToRect(g.Pieces.Feet)

	return g
}
