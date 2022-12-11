package rucksack

import (
	"bufio"
	"strings"
	"testing"
)

var sampleInput = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`

func TestRucksack(t *testing.T) {
	items := []byte("abac")
	r := NewRucksack(items[:2], items[2:])
	d := r.FindDup()
	if d[0] != byte('a') {
		t.Fail()
	}
}

func TestSolutionForPart1(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(sampleInput))
	sum := Solution(s, 0)

	if sum != 157 {
		t.Fatalf("sum of priority: %d\n", sum)
	}
}

func TestSolutionForPart2(t *testing.T) {
	s := bufio.NewScanner(strings.NewReader(sampleInput))
	sum := Solution(s, 3)

	if sum != 70 {
		t.Fatalf("sum of priority: %d\n", sum)
	}
}
