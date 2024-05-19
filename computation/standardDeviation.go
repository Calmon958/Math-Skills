package compute

import "math"

func StandardDeviation(num []float64) float64 {
	standard := Variance(num)
	standardDeviation := math.Sqrt(standard)
	return standardDeviation
}
