package main

import (
	"math"
	"strconv"
	"fmt"
)


func toFixed (f float64) float64 {
	val, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
	return val
}

func toFixed2 (f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func main () {
	var num = 123.0
	var num2 = 123.4567890

	fmt.Println(toFixed(num))
	fmt.Println(toFixed(num2))

	fmt.Println(toFixed2(num, 2))
	fmt.Println(toFixed2(num2, 2))
}