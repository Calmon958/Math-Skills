package compute

import (
	"sort"
)
//Takes a slice of floats and gives the median(middle) number
func Median(content []float64) float64 {
	dataFloat := sort.Float64Slice(content)
	answer := 0.0
	
	mid := float64(len(dataFloat) / 2)
	if len(dataFloat)%2 == 0 {
		answer = (float64(mid) + float64(mid-1)) / 2
	} else {
		answer = mid
	}

	return answer
}
