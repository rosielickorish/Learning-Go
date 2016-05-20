package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z:= float64(1)
	result := 0.0
	for {
		result = z - ((z*z-x)/2*z)
		
		if math.Abs(z- result)<0.001 {
			break
		}
		z = result
	}
	return math.Abs(z)

}

func main() {
	fmt.Println(Sqrt(2))	
	fmt.Println(Sqrt(4))
}

