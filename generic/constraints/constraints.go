package main

import (
	"fmt"

	. "golang.org/x/exp/constraints"
)

func min[T Ordered](x, y T) T {
	if x < y {
		return x
	}

	return y
}

func main() {
	x1 := 1
	x2 := 2
	y1 := 1.1
	y2 := 2.2
	fmt.Println(min(x1, x2))
	fmt.Println(min[float64](y1, y2))
}
