package file

import (
	"bufio"
	"fmt"
	"os"
)

type ProcessFunc func(string)

func ProcessFile(filename string, process ProcessFunc) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open file: %w", err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		process(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err) 
	}

	return nil
}

func ParseFile (filename string, f func(string)) error {
  return ProcessFile(filename, func(line string) { f(line )})
}