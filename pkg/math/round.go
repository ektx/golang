package main

import (
	"fmt"
	"math"
)

func main() {
	// 四舍五入
	p := math.Round(10.5)
	fmt.Printf("%.1f\n", p) // => 11.0
	
	p2 := math.Round(10.4)
	fmt.Printf("%.1f\n", p2) // => 10.0

	n := math.Round(-10.5)
	fmt.Printf("%.1f\n", n) // => -11.0
	
	n2 := math.Round(-10.4)
	fmt.Printf("%.1f\n", n2) // => -10.0
}