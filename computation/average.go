package compute

func Average(floatNum []float64) float64 {

var sum float64 
sum = 0
for _, data := range floatNum {
	sum += data
}
a:= sum/float64(len(floatNum))
// fmt.Print(a)
return a
}
