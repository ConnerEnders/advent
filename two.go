package main

import (
	"fmt"
	"strings"
)

func two(part int) {
	mem := strings.Split(read("two.txt"), ",")
	if part == 1 {
		fmt.Println(compute(getInts(mem), 12, 2))
	} else {
		for noun := 0; noun < 100; noun++ {
			for verb := 0; verb < 100; verb++ {
				result := compute(getInts(mem), noun, verb)
				if result == 19690720 {
					fmt.Println(noun, verb)
				}
			}
		}
	}
}

func compute(mem []int, noun, verb int) int {
	mem[1] = noun
	mem[2] = verb
	for i := 0; mem[i] != 99; i += 4 {
		op := mem[i]
		p1 := mem[i+1]
		p2 := mem[i+2]
		store := mem[i+3]
		if p1 < 0 || p1 >= len(mem) ||
			p2 < 0 || p2 >= len(mem) ||
			store < 0 || store >= len(mem) {
			break
		}
		if op == 1 {
			mem[store] = mem[p1] + mem[p2]
		} else if op == 2 {
			mem[store] = mem[p1] * mem[p2]
		}
	}
	return mem[0]
}