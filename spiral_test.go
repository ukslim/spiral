package spiral

import (
	"testing"
)

const TOLERANCE = 0.0001

func TestSpiral(t *testing.T) {

	table := []struct {
		time      float64
		expectedX float64
		expectedY float64
	}{
		{0, 0, 0},
		{360, 0, 10},
		{90, 2.5, 0},
		{180, 0, -5},
		{270, -7.5, 0},
		{720, 0, 20},
	}
	spiral := &Spiral{
		Period: 360,
		Width:  10,
	}

	for _, test := range table {
		expected := &Coords{test.expectedX, test.expectedY}

		actual := spiral.CoordsAt(test.time)

		if !actual.Compare(expected, TOLERANCE) {
			t.Errorf("At %v, expected %v, got %v", test.time, expected, actual)
		}

	}
}
