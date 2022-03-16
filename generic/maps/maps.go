package main

import (
	"fmt"

	"golang.org/x/exp/maps"
)

func main() {
	map1 := map[string]int{"a": 1}
	map2 := map[string]int{"a": 1}
	fmt.Println(maps.Equal(map1, map2))

	ml := map[string]int{"a": 1}
	maps.Clear(ml)
}
