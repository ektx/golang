package main

import (
	"fmt"
	"math"
)

func main() {
	// 保留2位小数
	x := math.Abs(-2)
	fmt.Printf("%.2f\n", x) // => 2.00
	
	n := float64(2)
	fmt.Printf("%.2f\n", n) // => 2.00
	
	str := "2"
	fmt.Printf("%.2f\n", str) // Err %!f(string=2)
	
	// 保留一位小数
	y := math.Abs(2)
	fmt.Printf("%.1f\n", y) // => 2.0
}