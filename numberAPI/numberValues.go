package numberAPI

import "math"

var NumbersDict = map[string]float64{
	"e":       math.E,
	"pi":      math.Pi,
	"e^pi":    math.Pow(math.E, math.Pi),
	"2^sqrt2": math.Pow(2, math.Sqrt2),
}
