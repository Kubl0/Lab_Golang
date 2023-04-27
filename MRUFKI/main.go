package main

import (
	"flag"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {

	myApp := app.New()
	myWindow := myApp.NewWindow("MRUFKI")
	myCanvas := myWindow.Canvas()

	myWindow.Resize(fyne.NewSize(getSize()))
	myWindow.SetFixedSize(true)
	myCanvas.SetOnTypedKey(func(key *fyne.KeyEvent) {
		if key.Name == fyne.KeyEscape {
			myWindow.Close()
		}
	})

	myWindow.SetContent(grid)

	myWindow.ShowAndRun()
}

func getSize() (float32, float32) {
	x := flag.Int("w", 0, "x size")
	y := flag.Int("h", 0, "y size")
	flag.Parse()
	fx := float32(*x)
	fy := float32(*y)

	return fx, fy
}
