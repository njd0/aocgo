package day4

import (
	"aoc/m/file"
	"fmt"
)

var filename string = "./2024/day4/puzzle.txt"

func findXmasDiagonal (wordSearch []string, row int, col int) bool {
	width := len(wordSearch[row])
	height := len(wordSearch)

	if row == 0 || row == width - 1 || col == 0 || col == height - 1 {
		return false
	}
	
	left := col - 1
	right := col + 1
	up := row - 1
	down := row + 1

	leftDiagCount := map[string]int {
		"M": 1,
		"S": 1,
	}
	rightDiagCount := map[string]int {
		"M": 1,
		"S": 1,
	}

	leftDiagCount[string(wordSearch[up][left])]--
	leftDiagCount[string(wordSearch[down][right])]--

	rightDiagCount[string(wordSearch[down][left])]--
	rightDiagCount[string(wordSearch[up][right])]--

	return leftDiagCount["M"] == 0 && leftDiagCount["S"] == 0 && rightDiagCount["M"] == 0 && rightDiagCount["S"] == 0
}


var xmasMatch = "XMAS"
func findXmas (wordSearch []string, row int, col int) (matches int) {
	matches = 0
	width := len(wordSearch[row])
	height := len(wordSearch)
	word := ""

	// horz - left
	for i := col; i >= 0 && i > col - 4; i-- {
		word += string(wordSearch[row][i])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// horz - right
	for i := col; i < width && i < col + 4; i++ {
		word += string(wordSearch[row][i])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// vert - up
	for i := row; i >= 0 && i > row - 4; i-- {
		word += string(wordSearch[i][col])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// vert - down
	for i := row; i < height && i < row + 4; i++ {
		word += string(wordSearch[i][col])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// diagonal - left up
	for i := row; i >= 0 && i > row - 4; i-- {
		cI := col - (row - i)
		if (cI < 0) { break }
		word += string(wordSearch[i][cI])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// diagonal - left down
	for i := row; i < height && i < row + 4; i++ {
		cI := col - (i - row)
		if (cI < 0) { break }
		word += string(wordSearch[i][cI])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// diagonal - right up
	for i := row; i >= 0 && i > row - 4; i-- {
			cI := col + (row - i)
			if (cI >= width) { break }
			word += string(wordSearch[i][cI])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	// diagonal - right down
	for i := row; i < height && i < row + 4; i++ {
		cI := col + (i - row)
		if (cI >= width) { break }
		word += string(wordSearch[i][cI])
	}

	if (word == xmasMatch) { matches++ }
	word = ""

	return
}

func Day4Part1() {
  var wordSearch []string

	err := file.ParseFile(filename, func(line string) {
    wordSearch = append(wordSearch, line)
  })

  if err != nil {
    return
  }

	// go through every position in search
	// if letter is x, search horz (left & right), vertical (up & down), diagonal
	// - that is, search all 8 directions around starting index
	total := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if string(wordSearch[i][j]) == "X" {
				total += findXmas(wordSearch, i, j)
			}
		}
	}

	fmt.Println("Part 1 Answer: ", total)
}

func Day4Part2 () {
	var wordSearch []string

	err := file.ParseFile(filename, func(line string) {
    wordSearch = append(wordSearch, line)
  })

  if err != nil {
    return
  }

	// Find an "A"
	// diagonally check in each direction that there is one "M" and one "S"
	// if both directions satisfy, count it
	total := 0
	for i := 0; i < len(wordSearch); i++ {
		for j := 0; j < len(wordSearch[i]); j++ {
			if string(wordSearch[i][j]) == "A" {
				if findXmasDiagonal(wordSearch, i, j) {
					total++
				}
			}
		}
	}

	fmt.Println("Part 2 Answer: ", total)
}