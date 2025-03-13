package solver

import (
	"math"
)

func SolveQuadEquation(a, b, c float64) (x1, x2 float64) {
	d := b*b - 4*a*c

	if d < 0 {
		return math.NaN(), math.NaN()
	}

	x1 = (-b + math.Sqrt(d)) / (2 * a)
	x2 = (-b - math.Sqrt(d)) / (2 * a)
	
	return
}