package main

import (
	"flag"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"math"
	"os"
	"spiral"
)

var mono = &color.Palette{
	color.RGBA{0x00, 0x00, 0x00, 0xff},
	color.RGBA{0xff, 0xff, 0xff, 0xff},
}

var width = flag.Float64("width", 5, "width of spiral path")
var period = flag.Float64("period", 360, "period of circle")
var infile = flag.String("in", "scratchy.png", "input file path")
var outfile = flag.String("out", "scratchy.spir", "input file path")

func main() {
	flag.Parse()

	reader, err := os.Open(*infile)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	writer, err := os.Create(*outfile)
	if err != nil {
		log.Fatal(err)
	}
	defer writer.Close()

	m, _, err := image.Decode(reader)

	if err != nil {
		log.Fatal(err)
	}

	sp := &spiral.Spiral{*width, *period}

	for i := 0; ; i++ {
		coords := sp.CoordsAt(float64(i))

		x := coords.X
		y := coords.Y

		if math.Abs(x) > 100 || math.Abs(y) > 100 {
			break
		}

		grayColor := mono.Convert(m.At(int(x)+100, int(y)+100))
		red, _, _, _ := grayColor.RGBA()

		if red > 100 {
			_, err = writer.Write([]byte{'1'})
		} else {
			_, err = writer.Write([]byte{'0'})
		}

		if err != nil {
			log.Fatal(err)
		}

	}

}
