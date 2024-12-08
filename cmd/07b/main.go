package main

import (
	"bufio"
	"fmt"
	"log"
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

func generateCombinations(length int, prefix string) [][]byte {
	if length == 0 {
		return [][]byte{[]byte(prefix)}
	}

	var combinations [][]byte
	for _, char := range []byte{'+', '*', '|'} {
		newPrefix := prefix + string(char)
		combinations = append(combinations, generateCombinations(length-1, newPrefix)...)
	}

	return combinations
}

func solve(equation equation_t) bool {
	// log.Print(equation.values)

	var bits [][]byte
	length := len(equation.values) - 1
	bits = append(bits, generateCombinations(length, "")...)

	for _, b := range bits {
		//log.Printf("%+v", string(b))
		res := 0
		calc := ""
		for i := 0; i < len(equation.values); i++ {
			if i == 0 {
				calc = fmt.Sprintf("%d", equation.values[i])
			} else {
				calc = fmt.Sprintf("%s%d", calc, equation.values[i])
			}
			if i < len(equation.values)-1 {
				calc = fmt.Sprintf("%s %s ", calc, string(b[i]))
			}
		}
		//log.Print(calc)
		res = performcalc(calc)

		// log.Printf("calc: %s = %d looking for %d", calc, res, equation.result)
		if res == equation.result {
			log.Printf("Found: %s = %d", calc, equation.result)
			return true
		}
	}
	log.Print("Not Found: ", equation.result, equation.values)
	return false
}

func performcalc(s string) int {
	cp := strings.Split(s, " ")
	e := len(cp) - 1
	r := mustAtoi(cp[0])
	i := 0
	for {
		//log.Print(i, e)
		if i == e {
			break
		}
		//log.Print(cp[i+1])
		switch cp[i+1] {
		case "+":
			//log.Print(r, "+", cp[i+2], (i + 2))
			r = r + mustAtoi(cp[i+2])
			i += 2
		case "*":
			//log.Print(r, "*", cp[i+2], (i + 2))
			r = r * mustAtoi(cp[i+2])
			i += 2
		case "|":
			//log.Print(r, "|", cp[i+2], (i + 2))
			r = mustAtoi(fmt.Sprintf("%d%s", r, cp[i+2]))
			i += 2
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
