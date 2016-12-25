package palette

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"

	"github.com/siuyin/dflt"
)

func Example() {
	const (
		xImg = 640
		yImg = 480

		xSwatch = 50
		ySwatch = 50

		xPad = 50
		yPad = 50
	)

	pl := Kelly22

	bgColorIdx, err := dflt.EnvInt("BG_COLOR_IDX", 0)
	if err != nil {
		log.Fatal(err)
	}

	m := image.NewRGBA(image.Rect(0, 0, xImg, yImg))
	draw.Draw(m, m.Bounds(), &image.Uniform{pl[bgColorIdx]}, image.ZP, draw.Src)
	i := 0
MAINLOOP:
	for k := 0; k < (yImg)/(ySwatch+yPad); k++ {
		for j := 0; j < (xImg)/(xSwatch+xPad); j++ {
			draw.Draw(m, image.Rect(10+j*50, 10+k*50, 10+(j+1)*50, 10+(k+1)*50).Add(image.Point{xPad * j, yPad * k}), &image.Uniform{pl[i]}, image.ZP, draw.Src)
			i++
			if i >= 22 {
				break MAINLOOP
			}
		}
	}

	f, _ := os.Create(fmt.Sprintf("img%d.png", bgColorIdx))
	png.Encode(f, m)
}
