package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// var palette6 = []color.Color{color.White, color.Black}
var palette6 = []color.Color{color.Black,
	color.RGBA{0xFF, 0x00, 0x00, 0xFF}, // Red
	color.RGBA{0xFF, 0x7F, 0x00, 0xFF}, // Orange
	color.RGBA{0xFF, 0xFF, 0x00, 0xFF}, // Yellow
	color.RGBA{0x00, 0xFF, 0x00, 0xFF}, // Green
	color.RGBA{0x00, 0x00, 0xFF, 0xFF}, // Blue
	color.RGBA{0x4B, 0x00, 0x82, 0xFF}, // Indigo
	color.RGBA{0x94, 0x00, 0xD3, 0xFF}, // Violet
}

func lissajous6(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette6)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			coloredIndex := uint8(rand.Intn(len(palette6)-1) + 1)                        // random index excluding black
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), coloredIndex) // random color from palette6 excluding black
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func main_1_6() {
	lissajous6(os.Stdout)
}
