package main

import (
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
		three()
	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func read(file string) []byte {
	input, err := ioutil.ReadFile(file)
	check(err)
	return input
}

func getInt(s string) int {
	i, err := strconv.Atoi(s)
	check(err)
	return i
}

func getInts(bytes [][]byte) []int {
	ints := make([]int, 0)
	for _, b := range bytes {
		ints = append(ints, getInt(string(b)))
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
