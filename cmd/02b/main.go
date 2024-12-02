package main

import (
	"log"
	"math"
	"os"
	"strings"
)

const fileName = "data.txt"

func main() {
	rows := loadFile(fileName)

	res := 0
	ores := 0
	for _, rl := range rows {

		// copy array
		re := make([]int, len(rl.level))

		for d := 0; d < rl.length; d++ {
			copy(re, rl.level)
			if d > 0 {
				// take one element out of the re array
				re = append(re[:d-1], re[d:]...)
				re = append(re, 0)
				log.Printf("%d %d %v", d-1, d, re)
			}
			rp := row_t{level: re, length: rl.length - 1}
			r := process(rp)
			if r == 1 {
				res += 1
				break
			}
		}

		msg := ""
		if ores != res {
			msg = "Safe"
		}
		log.Printf("Row: %v %s", rl, msg)
		ores = res
	}

	log.Printf("Total Safe rows: %d", res)
}

func process(row row_t) int {
	dir := ""

	for i := 0; i < row.length-1; i++ {
		res := checkLevel(row.level[i], row.level[i+1], dir)

		if res == "invalid" {
			return 0
		}

		if dir == "" {
			if dir != res && dir != "" {
				return 0
			}
			dir = res
		}
	}

	return 1
}

func checkLevel(l0 int, l1 int, dir string) string {
	var r string
	if l0 < l1 {
		if dir == "" || dir == "up" {
			r = "up"
		} else {
			return "invalid"
		}
	}
	if l0 > l1 {
		if dir == "" || dir == "down" {
			r = "down"
		} else {
			return "invalid"
		}
	}
	d := math.Abs(float64(l0) - float64(l1))
	if d == 0 || d > 3 {
		return "invalid"
	}
	return r
}

type row_t struct {
	length int
	level  []int
}

func loadFile(file string) []row_t {
	var rows []row_t

	// load file to memory
	d, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// Load the left and right column numbers
	for _, r := range strings.Split(string(d), "\n") {
		if r != "" {
			rl := row_t{level: make([]int, 8)}
			rls := strings.Split(r, " ")
			for i := 0; i < len(rls); i++ {
				rl.level[i] = mustAtoi(rls[i])
			}
			rl.length = len(rls)
			rows = append(rows, rl)
		}
	}

	return rows
}
