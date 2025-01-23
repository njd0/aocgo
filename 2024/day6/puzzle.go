package day6

import (
	"aoc/m/file"
	"aoc/m/vector"
	"fmt"
	"regexp"
	"strings"
)

var filename string = "./2024/day6/puzzle.txt"

type vec = vector.Vector

type Visit struct {
	pos vec
	dir int
}

var guardPos vec
var guardDir int
var startingGuardPos vec
var startingGuardCharacter string = "^"
var obstacleCharacter string = "#"
var walls = make(map[vec]bool)
var gameMap [][]string
var visited = make(map[Visit]bool)

var upDir = vec{0,-1}
var rightDir = vec{1,0}
var downDir = vec{0,1}
var leftDir = vec{-1,0}
var dirs = []vec {upDir, rightDir, downDir, leftDir} 

var part1Answer = make(map[vec]bool)
var part2Answer int = 0

func rotateGuardDir(dir int) int {
	dir++
	dir %= len(dirs)
	return dir
}

func resetGameState() {
	guardPos = startingGuardPos
	guardDir = 0
	visited = make(map[Visit]bool)
}

func isOnMap(v vec) bool {
	return v.X >= 0 && v.X < len(gameMap[0]) && v.Y >= 0 && v.Y < len(gameMap)
}

func Day6() {
	// setup game state related variables
	// - player can be global, obstacle char (start pos & dir)
	// walls and player should be vector based
	// create vector based math to improve readability

	// visited: [vector]bool
	// PART 1 ANSWER: count visited

	err := file.ParseFile(filename, func(line string) {
		// find where character is starting
		if strings.Contains(line, startingGuardCharacter) {
			indx := strings.Index(line, startingGuardCharacter)
			startingGuardPos = vec { indx, len(gameMap)}
		}

		// find all walls
		if strings.Contains(line, obstacleCharacter) {
			re := regexp.MustCompile(obstacleCharacter)
			for _, v := range re.FindAllStringIndex(line, -1) {
				walls[vec{v[0], len(gameMap)}] = true
			}
		}

		gameMap = append(gameMap, strings.Split(line, ""))
	})

	if err != nil {
		panic(err)
	}

	resetGameState()
	for {
		next := guardPos.Add(dirs[guardDir])
		
		if !isOnMap(next) {
			break
		}
		
		if _, exists := walls[next]; exists {
			guardDir = rotateGuardDir(guardDir)
			continue
		}

		visited[Visit { next, guardDir }] = true
		part1Answer[next] = true

		guardPos = next
	}

	// PART 2: How many infinite loops?
	// To consider pos as potential loop
	// 1. we must have already visited pos
	// 2. pos + new dir must eventually hit wall to be in loop
	newWalls := make(map[vec]bool)
	for visit := range visited {
		pos := visit.pos
		dir := visit.dir
		// try current pos as obstacle

		prev := pos.Sub(dirs[dir])
		if prev.Equal(startingGuardPos) { continue }
		dir = rotateGuardDir(dir)

		for {
			next := prev.Add(dirs[dir])
			
			if !isOnMap(next) { break }
			
			if _, exists := walls[next]; exists {
				newWalls[pos] = true
				break
			}

			prev = next
		}
	}

	// for each new wall to try
	for wallPos := range newWalls {
		resetGameState()
		// try and find infinite loop
		for {
			next := guardPos.Add(dirs[guardDir])
			
			if !isOnMap(next) { break }

			// wall
			if _, exists := walls[next]; exists || next.Equal(wallPos) {
				guardDir = rotateGuardDir(guardDir)
				continue
			}

			// infinite loop
			if _, exists := visited[Visit { next, guardDir }]; exists {
				part2Answer++
				break;
			}

			visited[Visit { next, guardDir }] = true

			guardPos = next
		}
	}

	fmt.Println("PART 1 ANSWER", len(part1Answer))
	fmt.Println("PART 2 ANSWER", part2Answer)
}
