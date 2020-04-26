package main

import (
	"fmt"
	"math"
)

func main() {
	// 返回2的1次方
	fmt.Printf("%.2f\n", math.Exp2(1)) // => 2
	// 返回2的-3次方
	fmt.Printf("%.2f\n", math.Exp2(-3)) // => 0.12
}
