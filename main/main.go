package main

import (
	"bufio"
	"flag"
	"log"
	"os"

	"github.com/chukonu/advent-of-code-2022/game"
	"github.com/chukonu/advent-of-code-2022/rucksack"
)

func main() {
	input := flag.String("input", "", "path to the input file")

	flag.Parse()

	if *input == "" {
		log.Println("no input")
		return
	}

	log.Printf("input: %s\n", *input)

	f, err := os.Open(*input)

	if err != nil {
		log.Fatalln(err)
	}

	s := bufio.NewScanner(f)

	score := rucksack.Solution(s, 3)

	log.Printf("Score: %d", score)
}

// func sum(list []int) (s int) {
// 	for i := 0; i < len(list); i++ {
// 		s += list[i]
// 	}

// 	return s
// }

// func insert(list []int, n int) {
// 	j := -1

// 	for i := 0; i < len(list); i++ {
// 		if list[i] < n {
// 			j = i
// 		}
// 	}

// 	if j < 0 {
// 		return
// 	}

// 	list[j] = n
// }

func day2(input *bufio.Scanner) int { return game.Solution(input) }

func day3(input *bufio.Scanner) {
	answer1 := rucksack.Solution(input, 0)
	log.Printf("sum of priorities: %d\n", answer1)

	answer2 := rucksack.Solution(input, 3)
	log.Printf("sum of priorities: %d\n", answer2)
}
