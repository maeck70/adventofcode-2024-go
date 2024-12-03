package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const fileName = "data.txt"

func main() {
	res := 0

	str := loadFile(fileName)
	re := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don\'t\(\)`)

	matches := re.FindAllStringSubmatch(str, -1)

	enabled := true
	for _, match := range matches {
		if len(match[1]) != 0 {
			num1, _ := strconv.Atoi(match[1])
			num2, _ := strconv.Atoi(match[2])
			r := num1 * num2
			if enabled {
				log.Printf("%s = %d", match[0], r)
				res += r
			}
		} else {
			if strings.Contains(match[0], "do()") {
				fmt.Println("--> do()")
				enabled = true
			} else {
				fmt.Println("--> don't()")
				enabled = false
			}
		}
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
