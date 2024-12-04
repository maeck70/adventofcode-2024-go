package main

import (
	"fmt"
	"log"
	"os"
)

const fileName = "data.txt"
const width = 140
const height = 140
const word = "XMAS"

// const fileName = "data_test.txt"
// const width = 10
// const height = 10
// const word = "XMAS"

var found = 0

type matrix_t [][]string

func main() {
	data := loadFile(fileName)
	log.Print(data)

	// Horizontal L2R
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			data.scan(x, y, word, 1, 0)
			data.scan(x, y, word, 0, 1)
			data.scan(x, y, word, -1, 0)
			data.scan(x, y, word, 0, -1)
			data.scan(x, y, word, 1, 1)
			data.scan(x, y, word, 1, -1)
			data.scan(x, y, word, -1, 1)
			data.scan(x, y, word, -1, -1)
		}
	}

	log.Printf("Found %s %d times", word, found)
}

func (data matrix_t) scan(x int, y int, w string, dx int, dy int) {
	if x >= width || y >= height || x < 0 || y < 0 {
		return
	}

	log.Printf("Check %s at %d,%d", string(w[0]), x, y)

	cc := data[y][x]
	cw := string(w[0])
	if cc == cw {
		log.Printf("Found %s at %d,%d", string(w[0]), x, y)
		if len(w) == 1 {
			log.Printf("Found %s at %d,%d", word, x, y)
			found++
		} else {
			data.scan(x+dx, y+dy, w[1:], dx, dy)
		}
	}
}

func loadFile(file string) matrix_t {
	// Setup runes slices
	var runes = make(matrix_t, height)
	for i := range runes {
		runes[i] = make([]string, width)
	}

	// load file to memory
	d, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// Load the data into a 2D slice of runes
	row := 0
	col := 0
	fmt.Printf("%03d : ", row)
	for _, c := range string(d) {
		if c == '\n' {
			row++
			col = 0
			fmt.Printf("\n%03d : ", row)
		} else {
			runes[row][col] = string(c)
			fmt.Print(string(c))
			col++
		}
	}
	fmt.Print("\n")

	return runes
}
