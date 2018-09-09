package pmg

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"os"
)

func EncodeImageToBase64(imagePath string) string {
	image, err := os.Open(imagePath)

	if err != nil {
		fmt.Println(err)
	}

	defer image.Close()

	// OpenFile to PNG
	decodedImage, err := png.Decode(image)

	// Encode the image to a base64 string
	var buffer bytes.Buffer
	png.Encode(&buffer, decodedImage)
	base64String := base64.StdEncoding.EncodeToString(buffer.Bytes())

	return base64String
}
