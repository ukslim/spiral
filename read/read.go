package main

import (
	"flag"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"log"
	"os"
	"spiral"
)

var infile = flag.String("in", "../write/scratchy.spir", "input file")
var width = flag.Float64("width", 5, "width of spiral path")
var period = flag.Float64("period", 360, "period of circle")
var outfile = flag.String("out", "out.png", "input file path")

func main() {
	flag.Parse()

	sfile, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		sfile.Close()
	}()

	sp := &spiral.Spiral{*width, *period}
	m := image.NewGray(image.Rect(-100, -100, 100, 100))
	white := color.Gray{255}
	draw.Draw(m, m.Bounds(), &image.Uniform{white}, image.ZP, draw.Src)
	for i := 0; ; i++ {
		b := make([]byte, 1)
		_, err = sfile.Read(b)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		coords := sp.CoordsAt(float64(i))
		if b[0] == '0' {
			m.Set(int(coords.X), int(coords.Y), color.Gray{0})
		}
	}

	f, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}

	if err := png.Encode(f, m); err != nil {
		f.Close()
		log.Fatal(err)
	}

	if err := f.Close(); err != nil {
		log.Fatal(err)
	}
}
