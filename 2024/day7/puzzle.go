package day7

import (
	"aoc/m/file"
	"fmt"
	"strconv"
	"strings"
)

var filename string = "./2024/day7/puzzle.txt"

type Input struct{
	target int
	set []int
}

func loadFile() []Input {
	inputs := make([]Input, 0)
	err := file.ParseFile(filename, func(line string) {
		var setup = strings.Split(line, ":")
		answer, err := strconv.Atoi(setup[0])
		if err != nil { panic(err) }
	
		values := make([]int, 0)
		for _, str := range strings.Fields(setup[1]) {
			value ,err := strconv.Atoi(str)
			if err != nil { panic(err) }
			values = append(values, value)
		}

		input := Input {
			target: answer,
			set: values,
		}

		inputs = append(inputs, input)
	})

	if err != nil {
		panic(err)
	}

	return inputs
}

func Day7Part1() {
	inputs := loadFile()

	ans := 0
	for _, input := range inputs {
		if dfs(input.target, input.set, []string {"+", "*"}) {
			ans += input.target
		}
	}

	fmt.Println("part1Answer", ans)
}

func Day7Part2() {
  inputs := loadFile()

	ans := 0
	for _, input := range inputs {
		if dfs(input.target, input.set, []string {"+", "*", "||"}) {
			ans += input.target
		}
	}

	fmt.Println("part1Answer", ans)
}

func dfs(target int, set []int, ops []string) bool {
	if len(set) == 1 {
		return target == set[0]
	}

	if set[0] > target {
		return false
	}

	for _, op := range ops {
		next := 0
		if op == "+" {
			next = set[0] + set[1]
		} else if op == "*" {
			next = set[0] * set[1]
		} else if op == "||" {
			str := fmt.Sprintf("%d%d", set[0], set[1])
			i, err := strconv.Atoi(str)
			if err != nil { panic(err) }
			next = i
		}

		if dfs(target, append([]int{next}, set[2:]...), ops) {
			return true
		}
	}

	return false
}