package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	switch getInt(os.Args[1]) {
	case 1:
		one()
	case 2:
		two()
	case 3:
		fmt.Println(sliceInt(1239487192, make([]int, 0)))
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(file string) string {
	input, err := ioutil.ReadFile(file)
	check(err)
	return string(input)
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

func sliceInt(n int, slice []int) []int {
	if n != 0 {
		i := n % 10
		slice = append([]int{i}, slice...)
		return sliceInt(n/10, slice)
	}
	return slice
}
