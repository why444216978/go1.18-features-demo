package main

import (
	"fmt"

	"golang.org/x/exp/slices"
)

func main() {
	slice1 := []int{1}
	slice2 := []int{1}
	fmt.Println(slices.Equal(slice1, slice2))

	slice3 := []int{10, 5, -3, 60, -8}
	fmt.Println(slices.IsSorted(slice3))
	slices.Sort(slice3)
	fmt.Println(slice3)
	fmt.Println(slices.IsSorted(slice3))
}
