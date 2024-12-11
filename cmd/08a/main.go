package main

import (
	"bufio"
	"log"
	"os"
)

const FREQS = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// const fileName = "data_test1.txt"
// const fileName = "data_test5.txt"
// const mw = 12
// const mh = 12

// const fileName = "data_test2.txt"
// const fileName = "data_test3.txt"
// const fileName = "data_test4.txt"
// const mw = 10
// const mh = 10

const fileName = "data.txt"
const mw = 50
const mh = 50

type map_t struct {
	w        int
	h        int
	antmap   [][]string
	freq     []freq_t
	freqs    []rune
	antinode []freq_t
}

type pos_t struct {
	x int
	y int
}

type freq_t struct {
	freq rune
	pos  pos_t
}

type freqpair_t struct {
	f1 pos_t
	f2 pos_t
}

func main() {
	m := map_t{w: mw, h: mh}
	m.loadFile(fileName)

	for _, r := range m.antmap {
		log.Print(r)
	}
	log.Print("---")

	for _, f := range m.freqs {
		for _, fp := range m.getPairs(f) {
			log.Printf("Freq pair (%s): %d,%d %d,%d", string(f), fp.f1.x, fp.f1.y, fp.f2.x, fp.f2.y)
			anPos := m.calculateAntiNode(f, fp.f1, fp.f2)
			if anPos.x >= 0 && anPos.x < mw && anPos.y >= 0 && anPos.y < mh {
				log.Printf("  - Antinode: %d,%d", anPos.x, anPos.y)
			}
		}
	}

	log.Print("---")
	log.Print("Number of antinodes: ", len(m.antinode))
}

// Calculate number of unique antinodes (non overlapping)
func (m map_t) uniqueAntiNodes() int {
	cnt := 0

	// For each antinode, check if it overlaps with any other antinode
	for _, an := range m.antinode {
		cnt += m.inAntiNode(an)
	}

	return cnt
}

func (m map_t) inAntiNode(an freq_t) int {
	for _, a := range m.antinode {
		if a.pos.x == an.pos.x && a.pos.y == an.pos.y {
			return 1
		}
	}
	return 0
}

func (m *map_t) calculateAntiNode(f rune, f1 pos_t, f2 pos_t) pos_t {
	// calculate distance between f1 and f2
	dx := f2.x - f1.x
	dy := f2.y - f1.y

	// calculate antinode position
	anX := f2.x + dx
	anY := f2.y + dy

	m.addAntinode(f, pos_t{anX, anY})

	return pos_t{anX, anY}
}

func (m *map_t) addAntinode(f rune, an pos_t) {
	if !(an.x >= mw || an.y >= mh) {
		if an.x >= 0 && an.y >= 0 {
			if m.inAntiNode(freq_t{freq: f, pos: an}) == 0 {
				m.antinode = append(m.antinode, freq_t{freq: f, pos: an})
			}
		}
	}
}

// return an array of unique freq pairs for a specific freq
func (m map_t) getPairs(f rune) []freqpair_t {
	var pairs []freqpair_t
	for i := range m.freq {
		if m.freq[i].freq == f {
			for j := range m.freq {
				if m.freq[j].freq == f {
					if i != j {
						pairs = append(pairs, freqpair_t{f1: m.freq[i].pos, f2: m.freq[j].pos})
					}
				}
			}
		}
	}
	return pairs
}

func (m *map_t) loadFile(fileName string) {

	m.antmap = make([][]string, m.h)
	for i := range m.antmap {
		m.antmap[i] = make([]string, m.w)
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Panic("Error opening file:", err)
	}
	defer file.Close()

	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		for x := 0; x < mw; x++ {
			char := rune(line[x])
			if instr(char, FREQS) {
				m.freq = append(m.freq, freq_t{freq: char, pos: pos_t{x, y}})
				if !inarr(char, m.freqs) {
					m.freqs = append(m.freqs, char)
				}
			}
			m.antmap[y][x] = string(char)
		}
		y++
	}

	if err := scanner.Err(); err != nil {
		log.Panic("Error reading file:", err)
	}
}
