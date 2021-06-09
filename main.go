package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var c complex128
	var i float64
	var y float64
	for i = -10; i < 10; i+= 5 {
		for y = -10; y < 10; y+= 5 {
			c = complex(i / 10, y / 10)
			fmt.Println(c, mandelbrot(c))
		}
	}
}

func mandelbrot(c complex128) float64 {
	var z complex128
	var n complex128
	
	for cmplx.Abs(z) <= 2 && real(n) < 80 {
		z = z * z + c
		n += 1
	}
	return real(n)
}