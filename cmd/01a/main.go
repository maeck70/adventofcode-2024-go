package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fileName = "data.txt"

func main() {
	rowsL, rowsR := loadFile(fileName)

	sort.Ints(rowsL)
	sort.Ints(rowsR)

	dist := 0
	for i := 0; i < len(rowsL); i++ {
		d := abs(rowsR[i] - rowsL[i])
		log.Printf("Row %d: %d - %d = %d", i, rowsR[i], rowsL[i], d)
		dist += d
	}

	log.Printf("Total distance: %d", dist)
}

func loadFile(file string) ([]int, []int) {
	var rowsL []int
	var rowsR []int

	// load file to memory
	d, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// Load the left and right column numbers
	for _, r := range strings.Split(string(d), "\n") {
		if r != "" {
			rd := strings.Split(r, "   ")
			rowsL = append(rowsL, mustAtoi(rd[0]))
			rowsR = append(rowsR, mustAtoi(rd[1]))
		}
	}

	return rowsL, rowsR
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
