package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// const fileName = "data_test.txt"
// const rows = 9

const fileName = "data.txt"
const rows = 850

type equation_t struct {
	result int
	values []int
}

func main() {
	eqs := loadFile(fileName)
	total := 0

	for r := range eqs {
		if solve(eqs[r]) {
			total += eqs[r].result
		}
	}

	log.Printf("Total: %d", total)
}

func solve(equation equation_t) bool {
	log.Print(equation.values)

	d := int(math.Pow(2, float64(len(equation.values))-1))
	d2 := len(equation.values) - 1
	bits := make([]string, d2)

	for j := 0; j < d; j++ {
		for x := 0; x < d2; x++ {
			if j&(1<<x) != 0 {
				bits[x] = "+"
			} else {
				bits[x] = "*"
			}
		}
		log.Printf("bits: %v", bits)

		var res int
		var calc string
		vn := len(equation.values)
		switch vn {
		case 2:
			calc = fmt.Sprintf("%d %s %d", equation.values[0], bits[0], equation.values[1])
			res = performcalc(calc)
		case 3:
			calc = fmt.Sprintf("%d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2])
			res = performcalc(calc)
		case 4:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3])
			res = performcalc(calc)
		case 5:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3], equation.values[4])
			res = performcalc(calc)
		case 6:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3], equation.values[4], bits[4], equation.values[5])
			res = performcalc(calc)
		case 7:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3], equation.values[4], bits[4], equation.values[5], bits[5], equation.values[6])
			res = performcalc(calc)
		case 8:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3], equation.values[4], bits[4], equation.values[5], bits[5], equation.values[6], bits[6], equation.values[7])
			res = performcalc(calc)
		case 9:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3],
				equation.values[4], bits[4], equation.values[5], bits[5], equation.values[6], bits[6], equation.values[7], bits[7], equation.values[8])
			res = performcalc(calc)
		case 10:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3],
				equation.values[4], bits[4], equation.values[5], bits[5], equation.values[6], bits[6], equation.values[7], bits[7], equation.values[8], bits[8], equation.values[9])
			res = performcalc(calc)
		case 11:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3],
				equation.values[4], bits[4], equation.values[5], bits[5], equation.values[6], bits[6], equation.values[7], bits[7], equation.values[8], bits[8], equation.values[9], bits[9], equation.values[10])
			res = performcalc(calc)
		case 12:
			calc = fmt.Sprintf("%d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d %s %d", equation.values[0], bits[0], equation.values[1], bits[1], equation.values[2], bits[2], equation.values[3], bits[3],
				equation.values[4], bits[4], equation.values[5], bits[5], equation.values[6], bits[6], equation.values[7], bits[7], equation.values[8], bits[8], equation.values[9], bits[9], equation.values[10],
				bits[10], equation.values[11])
			res = performcalc(calc)
		default:
			log.Panic("Invalid number of values: ", vn)
		}

		log.Printf("calc: %s = %d looking for %d", calc, res, equation.result)
		if res == equation.result {
			log.Printf("Found: %s = %d", calc, equation.result)
			return true
		}
	}
	return false
}

func performcalc(s string) int {
	cp := strings.Split(s, " ")
	r := mustAtoi(cp[0])
	for i := 0; i < len(cp); i++ {
		if i%2 == 1 {
			if cp[i] == "+" {
				r += mustAtoi(cp[i+1])
			} else {
				r *= mustAtoi(cp[i+1])
			}
		}
	}
	return r
}

func loadFile(fileName string) []equation_t {
	eqs := make([]equation_t, rows)

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error opening file:", err)
	}
	defer file.Close()

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		s1 := strings.Split(line, ": ")
		s2 := strings.Split(s1[1], " ")
		eqs[y].result = mustAtoi(s1[0])
		for _, vs := range s2 {
			eqs[y].values = append(eqs[y].values, mustAtoi(vs))
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Panic("Error reading file:", err)
	}

	return eqs
}
