package main

import (
	"fmt"
	"log"
	"os"
)

const fileName = "data.txt"
const width = 140
const height = 140

// const fileName = "data_test.txt"
// const width = 10
// const height = 10

const word = "XMAS"

var found = 0

type matrix_t [][]string
type map_t [][]string

var maps []map_t

func init() {
	maps = make([]map_t, 4)
	maps[0] = map_t{
		{"M", ".", "M"},
		{".", "A", "."},
		{"S", ".", "S"},
	}
	maps[1] = map_t{
		{"M", ".", "S"},
		{".", "A", "."},
		{"M", ".", "S"},
	}
	maps[2] = map_t{
		{"S", ".", "S"},
		{".", "A", "."},
		{"M", ".", "M"},
	}
	maps[3] = map_t{
		{"S", ".", "M"},
		{".", "A", "."},
		{"S", ".", "M"},
	}
}

func main() {
	data := loadFile(fileName)
	log.Print(data)

	for y := 0; y < height-2; y++ {
		for x := 0; x < width-2; x++ {
			for _, m := range maps {
				data.scan(x, y, m)
			}
		}
	}
	log.Printf("Found %s %d times", word, found)
}

func (data matrix_t) scan(x int, y int, m map_t) {
	log.Printf("Check %s at %d,%d", m[0], x, y)
	if m[0][0] == data[y][x] {
		data.scanMap(x, y, m)
	}
}

func (data matrix_t) scanMap(x int, y int, m map_t) {
	for ym := 0; ym < len(m); ym++ {
		for xm := 0; xm < len(m[ym]); xm++ {
			if m[ym][xm] != "." {
				if m[ym][xm] != data[y+ym][x+xm] {
					return
				}
			}
		}
	}
	log.Printf("Found map %v at %d,%d", m, x, y)
	found++
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
