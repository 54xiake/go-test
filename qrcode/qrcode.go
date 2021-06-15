package main

import (
	"github.com/skip2/go-qrcode"
	"image/color"
	"log"
)

func main() {
	var png []byte
	png, err := qrcode.Encode("https://example.org", qrcode.Medium, 256)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(png)

	err = qrcode.WriteFile("https://example.org", qrcode.Medium, 256, "./qr1.png")
	if err != nil {
		log.Fatal(err)
	}

	err = qrcode.WriteColorFile("https://example.org", qrcode.Medium, 256, color.Black, color.White, "./qr2.png")
	if err != nil {
		log.Fatal(err)
	}

}
