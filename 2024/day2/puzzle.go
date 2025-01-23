package day2

import (
	"aoc/m/file"
	"aoc/m/math"
	"fmt"
	"strconv"
	"strings"
)

var filename string = "./2024/day2/puzzle.txt"

func countValid(row []int) bool {
  // rules for valid #'s are all increasing or all decreasing
  // all adjecent #'s differ >= 1 && <= 3
  if (len(row) < 2){ return false}

  var incOrDec bool = row[0] < row[1]
  for i := 1; i < len(row); i++ {
    var dist = math.AbsInt(row[i - 1], row[i])
    if (incOrDec != (row[i-1] < row[i])) {return false}
    if (dist < 1 || dist > 3) {return false}
  }

  return true
}

func Day2Part1() {
  var inputs [][]int

  err := file.ParseFile(filename, func(line string) {
    list := strings.Fields(line)

    var nums = make([]int, len(list))
    for i := 0; i < len(list); i++ {
      integer, err := strconv.Atoi(list[i])
      if (err == nil) {
        nums[i] = integer
      }
    }

    inputs = append(inputs, nums)
  })

  if (err != nil) {
    panic(err)
  }

  total := 0
  for i := 0; i < len(inputs); i++ {
    if isValid(inputs[i]) { total++ }
  }

  fmt.Println("Part 1 Answer: ", total)
}

func isRemoveOneValid(row []int) bool {
	for i := 0; i < len(row); i++ {
		tmp := append([]int{}, row[:i]...)
		tmp = append(tmp, row[i+1:]...)
		if (isValid(tmp)) {
			return true
		}
	}

	return false
}

func isValid(row []int) bool {
	// rules for valid #'s are all increasing or all decreasing
  // all adjecent #'s differ >= 1 && <= 3
	isInc := true
	isDec := true
  for i := 1; i < len(row); i++ {
    var dist = math.AbsInt(row[i - 1], row[i])
    if (dist < 1 || dist > 3) {
			return false
		}

		if (row[i] > row[i-1]) {
			isDec = false
		}
		if (row[i] < row[i-1]) {
			isInc = false
		}
  }

  return isInc || isDec
}

func Day2Part2() {
	var inputs [][]int

  err := file.ParseFile(filename, func(line string) {
    list := strings.Fields(line)

    var nums = make([]int, len(list))
    for i := 0; i < len(list); i++ {
      integer, err := strconv.Atoi(list[i])
      if (err == nil) {
        nums[i] = integer
      }
    }

    inputs = append(inputs, nums)
  })

  if (err != nil) {
    panic(err)
  }

	total := 0
  for i := 0; i < len(inputs); i++ {
		if isValid(inputs[i]) || isRemoveOneValid(inputs[i]) {
			total++
		}
  }

	fmt.Println("Part 2 Answer: ", total)
}