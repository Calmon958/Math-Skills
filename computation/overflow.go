package compute

import (
"fmt"
"math"
)

func Overflow(float float64) int64 {
var ans int64
if float > float64(math.MaxInt) || float < float64(math.MinInt){
fmt.Printf("Value out of range")
return 0
}

ans = int64(math.Round(float))
return ans
}

