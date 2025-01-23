package day3

import (
	"aoc/m/file"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var filename string = "./2024/day3/puzzle.txt"

func countMul (instructions []string) int {
  total := 0
  for _, instruct := range instructions {
    re := regexp.MustCompile(`\d{1,3}`)
    digits := re.FindAllString(instruct, -1)
    left, _ := strconv.Atoi(digits[0])
    right, _ := strconv.Atoi(digits[1])
    total += left * right
  }

  return total
}

func getMulInstructions (str string) []string {
		re, err := regexp.Compile(`mul\(\d+,\d+\)`)
    if err != nil {
      fmt.Println("ERR: ", err)
    }
	  return re.FindAllString(str, -1)
}

func Day3Part1() {
  var instructions []string

	err := file.ParseFile(filename, func(line string) {
    instructions = append(instructions, getMulInstructions(line)...)
  })

  if err != nil {
    return
  }

  fmt.Println("Part 1 Answer: ", countMul(instructions))
}

func Day3Part2() {
	var inputString string
	err := file.ParseFile(filename, func(line string) {
		inputString += line
  })
	
	var instructions []string
	// each valid string starts after do, expect start
	chunks := strings.Split(inputString, "do()")
	for _, str := range chunks {
		index := strings.Index(str, "don't")
		end := len(str)
		if index != -1 {
			end = index
		}

		tmp := str[:end]
		
		instructions = append(instructions, getMulInstructions(tmp)...)
	}

  if err != nil {
    return
  }

  fmt.Println("Part 2 Answer: ", countMul(instructions))
}

