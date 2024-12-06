package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// const fileName = "data_test.txt"
// const width = 10
// const height = 10

const fileName = "data.txt"
const width = 130
const height = 130

type map_t [width][height]string

type dir_t struct {
	h int
	v int
}
type pos_t struct {
	x int
	y int
}

func main() {
	sm := loadFile(fileName)
	total := 0

	for ox := 0; ox < width; ox++ {
		for oy := 0; oy < height; oy++ {

			d := dir_t{0, -1}

			// Make a copy of the map and insert an obstruction
			var m map_t
			mapCopy(&m, &sm)
			// log.Printf("Obstruction at %d, %d", ox, oy)
			oob := false

			p := findPos(&m)
			m[ox][oy] = "O"

			var c map_t
			setMap(&c)

			if !(p.x == ox && p.y == oy) {
				for b := 0; b < 10000; b++ {
					// check next position
					switch m[p.x+d.h][p.y+d.v] {
					case "#", "O":
						rotate(&d)
					case ".", "^":
						move(&p, d, &c)
						// log.Print("Count Map:")
						// printMap(c)
					default:
						log.Fatalf("Invalid character: %s", m[p.x+d.h][p.y+d.v])
					}

					// Check if out of bounds
					if p.x+d.h == -1 || p.x+d.h == width || p.y+d.v == -1 || p.y+d.v == height {
						c[p.x][p.y] = "*"
						// log.Printf("Out of bounds at %d, %d", p.x+d.h, p.y+d.v)
						oob = true
						break
					}
				}
				if !oob {
					log.Printf(">>> Found looping path at %d %d", ox, oy)
					total++

					log.Printf("Iteration %d, %d", ox, oy)
					// printMap(m)
					// fmt.Println("---")
					printMap(c)
					fmt.Println("---")
					// fmt.Print("Press 'Enter' to continue...")
					// bufio.NewReader(os.Stdin).ReadBytes('\n')
				}
			}
		}
	}
	log.Printf("Total: %d", total)
}

func mapCopy(dst *map_t, src *map_t) {
	for y := range src {
		for x := range src[y] {
			dst[x][y] = src[x][y]
		}
	}
}

func move(p *pos_t, d dir_t, c *map_t) {
	c[p.x][p.y] = "*"
	p.x += d.h
	p.y += d.v
}

func rotate(d *dir_t) {
	if d.h == 0 {
		if d.v == -1 {
			d.h = 1
			d.v = 0
		} else {
			d.h = -1
			d.v = 0
		}
	} else {
		if d.h == -1 {
			d.h = 0
			d.v = -1
		} else {
			d.h = 0
			d.v = 1
		}
	}
}

func findPos(m *map_t) pos_t {
	for y := range m {
		for x := range m[y] {
			if m[x][y] == "^" {
				// m[y][x] = "."
				// log.Printf("Found starting position at %d, %d", x, y)
				return pos_t{x, y}
			}
		}
	}
	log.Fatal("Starting position not found")
	return pos_t{-1, -1}
}

func loadFile(fileName string) map_t {

	// initialize 2d map
	var m map_t

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error opening file:", err)
	}
	defer file.Close()

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x, c := range line {
			m[x][y] = string(c)
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Panic("Error reading file:", err)
	}

	return m
}

func printMap(m map_t) {
	for y := range m {
		for x := range m[y] {
			fmt.Printf("%s", m[x][y])
		}
		fmt.Print("\n")
	}
}

func setMap(m *map_t) {
	for y := range m {
		for x := range m[y] {
			m[x][y] = "."
		}
	}
}
