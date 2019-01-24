package main

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"log"
	"os"
)


type Version int8

func main() {
	fmt.Println("Hello QR Code")

	file, err := os.Create("qrcode.png")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = GenerateQRCode(file ,"555-2368", Version(1))
	if err != nil {
		log.Fatal(err)
	}

}

func (v Version) PatterSize() int {
	return 4*int(v) + 17
}

func GenerateQRCode(w io.Writer,code string, version Version) error {
	size := version.PatterSize()
	img := image.NewRGBA(image.Rect(0, 0, size, size))
	return png.Encode(w, img)
}

