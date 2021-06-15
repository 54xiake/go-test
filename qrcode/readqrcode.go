package main

import (
	"github.com/tuotoo/qrcode"
	"log"
	"os"
)

func main() {
	fi, err := os.Open("qr2.png")
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer fi.Close()
	qrmatrix, err := qrcode.Decode(fi)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Println(qrmatrix.Content)
}
