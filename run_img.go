package main

import (
	"image"
	"io/ioutil"
	"log"

	"github.com/liamgens/PimpMyGeorge/pmg"
)

func readBase64File(fileName string) string {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("Failed to open: %s", err)
	}

	return string(b)
}

func main() {
	george64 := pmg.EncodeImageToBase64("georges/georgeforeman/georgeforeman3.png")
	clothing64 := pmg.EncodeImageToBase64("clothing/hats/comptonflatbrim.png")

	location := image.Rect(200, -3, 315, 65)

	george := pmg.TempGeorge{george64, []pmg.TempBling{pmg.TempBling{clothing64, location}}}

	pmg.CreateBlingImage(george, "bling-image-out.png")
}
