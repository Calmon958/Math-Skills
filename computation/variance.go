package compute

import "math"

func Variance(floatNum []float64) float64 {
var sum float64
average := Average(floatNum)


for _, digit := range floatNum {
digit = math.Pow((digit - average), 2)
sum+=digit
}

standard:= sum/ float64(len(floatNum))

return standard

}