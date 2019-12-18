package main

import (
	"bytes"
	"fmt"
)

type point struct {
	X, Y int
}

func three() {
	wires := bytes.Split(read("three.txt"), []byte("\n"))
	p1 := getPath(bytes.Split(wires[0], []byte(",")))
	p2 := getPath(bytes.Split(wires[1], []byte(",")))
	// part 1
	tangles, shortest := getTangles(p1, p2)
	closest := closestTangle(tangles)
	fmt.Println(closest)
	// part 2
	fmt.Println(shortest)
}

func getPath(wire [][]byte) []point {
	location := new(point)
	path := make([]point, 0)

	for _, w := range wire {
		length := getInt(string(w[1:]))
		direction := new(point)
		switch string(w[0]) {
		case "U":
			direction.Y = 1
		case "R":
			direction.X = 1
		case "D":
			direction.Y = -1
		case "L":
			direction.X = -1
		}
		for i := 0; i < length; i++ {
			location.X += direction.X
			location.Y += direction.Y
			path = append(path, *location)
		}
	}

	return path
}

func getTangles(path1, path2 []point) ([]point, int) {
	tangles := make([]point, 0)
	shortest := -1
	for i, p1 := range path1 {
		for j, p2 := range path2 {
			if p1.X == p2.X && p1.Y == p2.Y {
				tangles = append(tangles, p1)
				distance := i + 1 + j + 1
				if distance < shortest || shortest == -1 {
					shortest = distance
				}
			}
		}
	}
	return tangles, shortest
}

func closestTangle(tangles []point) int {
	closest := -1
	for _, t := range tangles {
		distance := absInt(t.X) + absInt(t.Y)
		if distance < closest || closest == -1 {
			closest = distance
		}
	}
	return closest
}
