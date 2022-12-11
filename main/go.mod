module github.com/chukonu/advent-of-code-2022

go 1.17

replace github.com/chukonu/advent-of-code-2022/game => ../game

require (
	github.com/chukonu/advent-of-code-2022/game v0.0.0-00010101000000-000000000000
	github.com/chukonu/advent-of-code-2022/rucksack v0.0.0-00010101000000-000000000000
)

require golang.org/x/exp v0.0.0-20221205204356-47842c84f3db // indirect

replace github.com/chukonu/advent-of-code-2022/rucksack => ../rucksack
