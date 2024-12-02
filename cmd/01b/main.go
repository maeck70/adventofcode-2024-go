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

	sort.Ints(rowsR)

	result := 0
	for i := 0; i < len(rowsL); i++ {
		ss := findRepeats(rowsL[i], rowsR)
		result += rowsL[i] * ss
		log.Printf("Row %d (%d %d), %d repeats, score %d", i, rowsL[i], rowsR[i], ss, rowsL[i]*ss)
	}

	log.Printf("Total similarity score : %d", result)
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

func findRepeats(i int, arr []int) int {
	o := 0
	p := 0
	for _, a := range arr {
		if p > a {
			break
		}
		if a == i {
			o += 1
		}
		p = a
	}

	return o
}
