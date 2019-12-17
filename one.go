package main

import (
	"bytes"
	"fmt"
)

func one() {
	input := bytes.Split(read("one.txt"), []byte("\n"))
	modules := getInts(input)
	var total1, total2 int
	for _, mass := range modules {
		var fuel int
		// part 1
		total1 += getFuel(mass)
		// part 2
		for {
			mass = getFuel(mass)
			if mass <= 0 {
				break
			}
			fuel += mass
		}
		total2 += fuel
	}
	fmt.Println(total1)
	fmt.Println(total2)
}

func getFuel(mass int) int {
	return mass/3 - 2
}
