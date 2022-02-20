package square

import "math"

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

type sidesType int

const SidesTriangle sidesType = 3
const SidesSquare sidesType = 4
const SidesCircle sidesType = 0

func CalcSquare(sideLen float64, sidesNum sidesType) float64 {
	square := 0.0
	switch sidesNum {
	case SidesCircle:
		square = math.Pi * sideLen * sideLen
	case SidesSquare:
		square = sideLen * sideLen
	case SidesTriangle:
		square = (math.Sqrt(3) / 4) * math.Pow(sideLen, 2)
	}
	return square
}
