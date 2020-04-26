package main

import (
	"fmt"
	"math"
)

func main() {
	// 向上取整
	c := math.Ceil(1.49)
	fmt.Printf("%.1f\n", c) // => 2.0
	
	d := math.Ceil(1.01)
	fmt.Printf("%.1f\n", d) // => 2.0
	
	e := math.Ceil(0.01)
	fmt.Printf("%.1f\n", e) // => 1.0
}