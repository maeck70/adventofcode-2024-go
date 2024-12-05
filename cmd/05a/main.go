package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// const fileName = "data_test.txt"
const fileName = "data.txt"

type rule_t struct {
	pageBefore int
	pageAfter  int
}

type update_t []int

func main() {
	rules, updates := loadFile(fileName)
	log.Print("Rules: ", rules)
	log.Print("Updates: ", updates)

	total := 0

	for upd := range updates {
		log.Print("Update: ", updates[upd])
		if check(rules, updates[upd]) {
			k := updates[upd][len(updates[upd])/2]
			log.Printf("Valid: key is %d", k)
			total += k
		} else {
			log.Print("Invalid")
		}
	}

	log.Printf("Total: %d", total)
}

func check(rules []rule_t, update update_t) bool {
	for _, rn := range rules {
		scan := -1
		for i, un := range update {
			if rn.pageBefore == un {
				scan = i
				break
			}
		}
		if scan != -1 {
			// Check before
			for _, unb := range update[:scan] {
				if rn.pageAfter == unb {
					return false
				}
			}
			// Check after
			for _, unb := range update[scan:] {
				if rn.pageAfter == unb {
					break
				}
			}
		}
	}
	return true
}

func loadFile(fileName string) ([]rule_t, []update_t) {
	var rules []rule_t
	var updates []update_t

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error opening file:", err)
	}
	defer file.Close()

	mode := "rules"
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		switch mode {
		case "rules":
			if line == "" {
				mode = "updates"
				continue
			}

			r := strings.Split(line, "|")
			rules = append(rules, rule_t{mustAtoi(r[0]), mustAtoi(r[1])})

		case "updates":
			pu := update_t{}
			for _, i := range strings.Split(line, ",") {
				pu = append(pu, mustAtoi(i))
			}
			updates = append(updates, pu)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Panic("Error reading file:", err)
	}

	return rules, updates
}

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
