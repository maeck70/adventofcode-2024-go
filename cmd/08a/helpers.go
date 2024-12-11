package main

import "strconv"

func mustAtoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func instr(s rune, a string) bool {
	if s == rune('.') {
		return false
	}

	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}

func inarr(s rune, a []rune) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
