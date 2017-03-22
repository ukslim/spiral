package spiral

import (
	"fmt"
	"math"
)

type Spiral struct {
	Width  float64
	Period float64
}

type Coords struct {
	X float64
	Y float64
}

func (c Coords) String() string {
	return fmt.Sprintf("[%0.002f, %0.002f]", c.X, c.Y)
}

func (c Coords) Compare(other *Coords, tolerance float64) bool {
	return math.Abs(c.X-other.X) < tolerance &&
		math.Abs(c.Y-other.Y) < tolerance
}

func (s Spiral) CoordsAt(time float64) *Coords {
	radius := time * s.Width / s.Period
	angle := math.Pi * 2 * (time / s.Period)
	x := math.Sin(angle) * radius
	y := math.Cos(angle) * radius
	return &Coords{x, y}
}
