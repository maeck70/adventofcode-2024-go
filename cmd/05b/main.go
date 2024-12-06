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

var skip []int

func main() {
	rules, updates := loadFile(fileName)
	total := 0

	skip = make([]int, 0)

	// Make skip list for updates that are valid
	for upd := range updates {
		if check(rules, updates[upd]) {
			skip = append(skip, upd)
		}
	}

	// Find invalid pages and swap the rule before/after
	// Do this until no invalid pages in the updates are found
	for {
		invalids := 0
		for upd := range updates {
			if !inSkip(upd) && !check(rules, updates[upd]) {
				log.Print(updates[upd])
				invalids += findInvalidPage(rules, updates[upd])
				log.Print(updates[upd])
				log.Print(" ")
			}
		}
		log.Printf("Invalids: %d", invalids)
		if invalids == 0 {
			break
		}
		log.Print("--- AGAIN ---")
	}

	// Count Total
	for upd := range updates {
		if !inSkip(upd) {
			k := updates[upd][len(updates[upd])/2]
			total += k
		}
	}

	log.Printf("Total: %d", total)
}

func inSkip(i int) bool {
	for _, s := range skip {
		if i == s {
			return true
		}
	}
	return false
}

func findInvalidPage(rules []rule_t, update update_t) int {
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
			for i, unb := range update[:scan] {
				if rn.pageAfter == unb {
					// Swap pages
					update[changeAfter(update, rn.pageBefore)] = rn.pageAfter
					update[i] = rn.pageBefore
					log.Printf("Invalid page: %d", unb)
					return 1
				}
			}
		}
	}
	return 0
}

func changeAfter(update update_t, find int) int {
	for i, un := range update {
		if un == find {
			return i
		}
	}
	log.Fatalf("Page not found: %d", find)
	return 0
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
