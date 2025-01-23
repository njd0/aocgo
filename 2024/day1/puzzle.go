package day1

import (
	"aoc/m/file"
	"aoc/m/math"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var filename string = "./2024/day1/puzzle.txt"

func Day1Part1() {
    var list1 []int
    var list2 []int

		err := file.ParseFile(filename, func(line string) {
      lists := strings.Fields(line)
      i1, err1 := strconv.Atoi(lists[0])
      if (err1 == nil) {
        list1 = append(list1, i1)
      }

      i2, err2 := strconv.Atoi(lists[1])
      if (err2 == nil) {
        list2 = append(list2, i2)
      }
	  })

    if (err != nil) {
      return
    }

    sort.Ints(list1)
    sort.Ints(list2)

    if (len(list1) != len(list2)) {
      return
    }

    totalDist := 0
    for i := 0; i < len(list1); i++ {
      totalDist += math.AbsInt(list1[i], list2[i])
    }

    fmt.Println("Part 1 Answer: ", totalDist)
}

func Day1Part2() {
  var list1 []string
  var list2 []string

  err := file.ParseFile(filename, func(line string) {
    lists := strings.Fields(line)
    list1 = append(list1, lists[0])
    list2 = append(list2, lists[1])
  })

  if (err != nil) {
    return
  }

  mapList2 := make(map[string]int)

  for i := 0; i < len(list2); i++ {
    mapList2[list2[i]]++
  }
  
  distance := 0
  for i := 0; i < len(list1); i++ {
    if (mapList2[list1[i]] > 0) {
      i10, err := strconv.Atoi(list1[i])
      if (err != nil) {continue}
      distance += (i10 * mapList2[list1[i]])
    } 
  }

  fmt.Println("Part 2 Answer: ", distance)
}