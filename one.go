package main

import (
	"fmt"
	"strings"
)

func one(part int) {
	input := strings.Split(read("one.txt"), "\n")
	modules := getInts(input)
	totalFuel := 0
	for _, mass := range modules {
		fuel := 0
		if part == 1 {
			totalFuel += getFuel(mass)
		} else {
			for {
				mass = getFuel(mass)
				if mass <= 0 {
					break
				}
				fuel += mass
			}
			totalFuel += fuel
		}
	}
	fmt.Println(totalFuel)
}

func getFuel(mass int) int {
	return mass/3 - 2
}
