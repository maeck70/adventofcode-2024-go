package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
)

const fileName = "data.txt"

func main() {
	res := 0

	str := loadFile(fileName)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	matches := re.FindAllStringSubmatch(str, -1)

	for _, match := range matches {
		num1, _ := strconv.Atoi(match[1])
		num2, _ := strconv.Atoi(match[2])
		r := num1 * num2
		log.Printf("%s = %d", match[0], r)
		res += r
	}

	log.Printf("Calculated total: %d", res)
}

func loadFile(file string) string {

	// load file to memory
	d, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return string(d)
}
