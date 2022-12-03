package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(calc(10, 10, 2.5))
}

func calc(x int16, y uint8, z float32) float32 {
	res := float64(x)*2 + math.Pow(float64(y), 2) - 3/float64(z)
	return float32(res)
}
