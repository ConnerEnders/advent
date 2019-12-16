package main

import (
	"flag"
	"io/ioutil"
	"strconv"
)

func main() {
	day := flag.Int("d", 2, "Advent day")
	part := flag.Int("p", 2, "Advent day part")
	flag.Parse()

	switch *day {
	case 1:
		one(*part)
	case 2:
		two(*part)
	}
}

func read(file string) string {
	input, err := ioutil.ReadFile(file)
	check(err)
	return string(input)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func getInts(strings []string) []int {
	ints := make([]int, 0)
	for _, s := range strings {
		ints = append(ints, getInt(s))
	}
	return ints
}
