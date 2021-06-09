package main

import (
	"math/cmplx"
	"image"
	"image/color"
	"image/png"
	"os"
)


const MAX_ITER = 80

const IMG_WIDTH = 600
const IMG_HEIGHT = 400

const RE_START = -2
const RE_END = 1
const IM_START = -1
const IM_END = 1




func main() {
	createImage()

}

func createImage() {
	upLeft := image.Point{0,0}
	lowRight := image.Point{IMG_WIDTH, IMG_HEIGHT}
	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel.
	for x := 0; x < IMG_WIDTH; x++ {
		for y := 0; y < IMG_HEIGHT; y++ {
			c := complex(float64(RE_START) + (float64(x) / float64(IMG_WIDTH)) * (float64(RE_END) - float64(RE_START)), 
						float64(IM_START) + (float64(y) / float64(IMG_HEIGHT)) * (float64(IM_END) - float64(IM_START)))
			m := mandelbrot(c, MAX_ITER)
			colorvalue := 255 - int(m * 255 / MAX_ITER)
			color := color.RGBA{uint8(colorvalue), uint8(colorvalue), uint8(colorvalue), 0xff}
			img.Set(x, y, color)
		}
}

// Encode as PNG.
f, _ := os.Create("image.png")
png.Encode(f, img)

}

func mandelbrot(c complex128, maxIter float64) float64 {
	var z complex128
	var n complex128
	
	for cmplx.Abs(z) <= 2 && real(n) < maxIter {
		z = z * z + c
		n += 1
	}
	return real(n)
}