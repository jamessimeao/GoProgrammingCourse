package main

import "fmt"

func main() {
	// A map
	var mp map[string]int = map[string]int{"a": 1}
	fmt.Println(mp)

	var anotherMap map[string][]int = map[string][]int{"b": {1, 2, 3}}
	fmt.Println(anotherMap)

	mp["b"] = 2
	mp["d"] = 4
	fmt.Println(mp)
	delete(mp, "d")
	fmt.Println(mp)

	value, ok := mp["b"]
	fmt.Printf("value = %v\n", value)
	if !ok {
		fmt.Println("mp doesn't have a value for key b")
	}

	var uintMap map[uint]uint = map[uint]uint{}

	var i int
	const n int = 100
	for i = 1; i <= n; i++ {
		for d := 1; d <= 5; d++ {
			if i%d == 0 {
				uintMap[uint(d)]++
			}
		}
	}
	fmt.Println(uintMap)
}
