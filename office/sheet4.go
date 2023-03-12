package main

import (
	"fmt"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	f, err := excelize.OpenFile("output/Book1.xlsx2")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Insert a picture.
	if err := f.AddPicture("Sheet1", "A2", "image.png", nil); err != nil {
		fmt.Println(err)
	}
	// Insert a picture to worksheet with scaling.
	if err := f.AddPicture("Sheet1", "D2", "image.jpg",
		&excelize.GraphicOptions{ScaleX: 0.5, ScaleY: 0.5}); err != nil {
		fmt.Println(err)
	}
	// Insert a picture offset in the cell with printing support.
	enable, disable := true, false
	if err := f.AddPicture("Sheet1", "H2", "image.gif",
		&excelize.GraphicOptions{
			PrintObject:     &enable,
			LockAspectRatio: false,
			OffsetX:         15,
			OffsetY:         10,
			Locked:          &disable,
		}); err != nil {
		fmt.Println(err)
	}
	// Save the spreadsheet with the origin path.
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
}
