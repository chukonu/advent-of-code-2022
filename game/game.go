package game

import (
	"bufio"
	"strings"
)

var shapes = map[string]shape{
	"A": {name: "Rock", score: 1, defeating: "C", defeatedBy: "B"},
	"B": {name: "Paper", score: 2, defeating: "A", defeatedBy: "C"},
	"C": {name: "Scissors", score: 3, defeating: "B", defeatedBy: "A"},
}

var scores = map[string]int{
	"A": 1,
	"B": 2,
	"C": 3,
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var outcomes = map[string]int{
	"X": 0,
	"Y": 3,
	"Z": 6,
}

type Game struct {
	score  int
	rounds int
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Score() int {
	return g.score
}

func (g *Game) Rounds() int {
	return g.rounds
}

func (g *Game) Play(shape1, shape2 string) {
	g.rounds++

	g.score += scores[shape2] + compare(shape1, shape2)
}

func (g *Game) Deduce(shape, outcome string) {
	g.rounds++

	g.score += outcomes[outcome] + deduce(shape, outcome)
}

func deduce(shape, outcome string) (s int) {
	switch outcome {
	case "X":
		s = shapes[shapes[shape].defeating].score
	case "Y":
		s = shapes[shape].score
	case "Z":
		s = shapes[shapes[shape].defeatedBy].score
	}

	return s
}

func compare(shape1, shape2 string) (s int) {
	switch d := scores[shape1] - scores[shape2]; d {
	case 0:
		s = 3

	case 1:
		fallthrough
	case -2:
		s = 0

	case 2:
		fallthrough
	case -1:
		s = 6
	}

	return s
}

func Solution(s *bufio.Scanner) int {
	g := NewGame()

	for s.Scan() {
		txt := s.Text()
		tok := strings.Split(txt, " ")
		g.Deduce(tok[0], tok[1])
	}

	if err := s.Err(); err != nil {
		panic(err)
	}

	return g.Score()
}
