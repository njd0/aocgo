package day5

import (
	"aoc/m/file"
	"fmt"
	"strconv"
	"strings"
)

var filename string = "./2024/day5/puzzle.txt"
var rulesFilename string = "./2024/day5/rules.txt"

func getKey(p1 string, p2 string) string {
	return fmt.Sprintf("%s|%s", p1, p2)
}

func parseInput() (map[string]bool, [][]string){
	rules := make(map[string]bool)
	err := file.ParseFile(rulesFilename, func(line string) {
   	rules[line] = true
  })

  if err != nil {
		panic(err)
  }

	pages := [][]string {}
	err = file.ParseFile(filename, func(line string) {
   	pages = append(pages, strings.Split(line, ","))
  })

  if err != nil {
		panic(err)
  }

	return rules, pages
}

func validateUpdates (rules map[string]bool, pages [][]string) (verified [][]string, unverified [][]string) {
	for _, pageNumbers := range pages {
		isValid := true
		for i := 1; i < len(pageNumbers); i++ {
			// check page number order is validated in rules
			key := getKey(pageNumbers[i-1], pageNumbers[i])
			if !rules[key] { isValid = false; break }
		}

		if isValid {
			verified = append(verified, pageNumbers)
		} else {
			unverified = append(unverified, pageNumbers)
		}
	}

	return
}

func countUpdateTotals(updates [][]string) (total int) {
	for _, nums := range updates {
		mid := int(len(nums) / 2)
		middleNum, _ := strconv.Atoi(string(nums[mid]))
		total += middleNum
	}
	return
}

func Day5Part1 () {
	rules, pages := parseInput()

	verified, _ := validateUpdates(rules, pages)
	total := countUpdateTotals(verified)

	fmt.Println("Part 1 Answer: ", total)
}

func Day5Part2 () {
	rules, pages := parseInput()

	_, unverified := validateUpdates(rules, pages)

	var reordered [][]string
	for _, pageNumbers := range unverified {
		for i := 0; i < len(pageNumbers); i++ {
			swapped := false
			for j := 0; j < len(pageNumbers) - i - 1; j++ {
				if !rules[getKey(pageNumbers[j], pageNumbers[j+1])] {
					pageNumbers[j], pageNumbers[j+1] = pageNumbers[j+1], pageNumbers[j]
					swapped = true
				}
			}
			if !swapped {
				break
			}
		}
		reordered = append(reordered, pageNumbers)
	}

	total := countUpdateTotals(reordered)

	fmt.Println("Part 2 Answer: ", total)
}
