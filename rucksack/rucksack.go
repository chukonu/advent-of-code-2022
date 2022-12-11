package rucksack

import (
	"bufio"

	"golang.org/x/exp/slices"
)

func Solution(s *bufio.Scanner, compSize int) (priorities int) {
	var lines [][]byte

	for s.Scan() {
		// clone because the underlying array may be overwritten later
		currLine := slices.Clone(s.Bytes())

		lines = append(lines, currLine)

		if len(lines) >= compSize {
			var r *Rucksack

			if compSize > 1 {
				r = NewRucksack(lines...)
			} else {
				l := len(currLine)
				r = NewRucksack(currLine[:l/2], currLine[l/2:])
			}

			dup := r.FindDup()

			priorities += normalize(dup)

			lines = make([][]byte, 0)
		}
	}

	return priorities
}

func normalize(b []byte) (sum int) {
	for _, v := range b {
		if i := int(v); i > 96 {
			sum += i - 96
		} else {
			sum += i - 38
		}
	}

	return
}

type Rucksack struct {
	compartments [][]byte
}

func (r *Rucksack) FindDup() (dup []byte) {
	for _, c := range r.compartments {
		if dup == nil {
			dup = c
			continue
		}

		var ldup []byte

		for _, d := range dup {
			if slices.Contains(c, d) && !slices.Contains(ldup, d) {
				ldup = append(ldup, d)
			}
		}

		dup = ldup
	}

	return
}

func NewRucksack(compartments ...[]byte) *Rucksack {
	return &Rucksack{compartments}
}
