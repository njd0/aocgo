package day8

import (
	"aoc/m/file"
	"aoc/m/vector"
	"fmt"
	"regexp"
)

var filename = "./2024/day8/puzzle.txt"
type vec = vector.Vector

func loadInputs() (map[string][]vec, int, int) {
	row := 0
	width := 0
	antennas := make(map[string][]vec)
	err := file.ParseFile(filename, func (line string) {
		if width == 0 {
			width = len(line)
		}

		re, err := regexp.Compile(`[\d+|\w+]`)
		if err != nil {
			return
		}

		matches := re.FindAllStringIndex(line, -1)

		for _, match := range matches {
			char := string(line[match[0]])
			antennas[char] = append(antennas[char], vec {match[0], row})
		}

		row++
	})

	if err != nil {
		panic(err)
	}

	return antennas, row, width
}

func inBounds(v vec, height int, width int) bool{
	return v.X >= 0 && v.X < width && v.Y >= 0 && v.Y < height
}

func Day8Part1() {
	antennas, row, width := loadInputs()

	antinodes := make(map[vec]bool)
	for _, vecs := range antennas {
		for i, v1 := range vecs {
			for j, v2 := range vecs {
				if (i == j) { continue }
				normalize := v1.Sub(v2)
				antinode := v1.Add(normalize)
				if inBounds(antinode, row, width) {
					antinodes[antinode] = true
				}
			}
		}
	}

	fmt.Println("Part 1 Answer", len(antinodes))
}

func Day8Part2 () {
	antennas, row, width := loadInputs()
	antinodes := make(map[vec]bool)

	for _, vecs := range antennas {
		for i, v1 := range vecs {
			antinodes[v1] = true
			for j, v2 := range vecs {
				if (i == j) { continue }
				normalize := v1.Sub(v2)
				antinode := v1.Add(normalize)
				// antinodes get created repeatedly
				for inBounds(antinode, row, width) {
					antinodes[antinode] = true
					antinode = antinode.Add(normalize)
				}
			}
		}
	}



	fmt.Println("Part 2 Answer", len(antinodes))
}