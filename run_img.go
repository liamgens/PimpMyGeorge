package main

import (
	"image"
	"io/ioutil"
	"log"

	"github.com/liamgens/Gim/pkg/pmg"
)

// TODO: Remove this.
func readBase64File(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to open: %s", err)
	}

	return string(b)
}

func main() {
	curiousGeorge := readBase64File("cg1.txt")
	mustache := readBase64File("mustache.txt")
	pinkShirt := readBase64File("pinkshirt.txt")

	george := pmg.TempGeorge{curiousGeorge, []pmg.TempBling{pmg.TempBling{mustache, image.Rect(190, 170, 290, 220)}, pmg.TempBling{pinkShirt, image.Rect(30, 300, 330, 600)}}}

	pmg.CreateBlingImage(george, "bling-image.png")
}
