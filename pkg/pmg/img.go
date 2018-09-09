package pmg

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/anthonynsimon/bild/transform"
)

// TODO: Remove this.
type TempGeorge struct {
	Image  string
	Blings []TempBling
}

// TODO: Remove this.
type TempBling struct {
	Image    string
	Location image.Rectangle
}

func CreateBlingImage(george TempGeorge, outputFileName string) error {
	georgeImg, err := base64AsPng(george.Image)
	if err != nil {
		return err
	}

	for _, bling := range george.Blings {
		err := addBling(georgeImg, bling)
		if err != nil {
			return err
		}
	}

	writePngToFile(georgeImg, outputFileName)

	return nil
}

func addBling(georgeImg *image.RGBA, bling TempBling) error {
	blingImg, err := base64AsPng(bling.Image)
	if err != nil {
		return err
	}

	dimensions := bling.Location.Max.Sub(bling.Location.Min)
	blingImg = transform.Resize(blingImg, dimensions.X, dimensions.Y, transform.Linear)

	draw.Draw(georgeImg, blingImg.Bounds().Add(bling.Location.Min), blingImg, image.ZP, draw.Over)

	return nil
}

func base64AsPng(imgBase64 string) (*image.RGBA, error) {
	imgBytes, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return &image.RGBA{}, err
	}

	imgImage, err := png.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return &image.RGBA{}, err
	}

	bounds := imgImage.Bounds()
	imgRGBA := image.NewRGBA(bounds)
	draw.Draw(imgRGBA, bounds, imgImage, image.ZP, draw.Src)

	return imgRGBA, nil
}

func writePngToFile(img *image.RGBA, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	png.Encode(file, img)

	return nil
}
