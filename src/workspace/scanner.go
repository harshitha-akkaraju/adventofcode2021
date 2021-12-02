package workspace

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

type Scanner struct {}

func NewScanner() *Scanner {
	return &Scanner{}
}

func (s *Scanner) ToIntArray(filePath string) []int {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	result := []int{}
	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, n)
	}

	return result
}

func (s *Scanner) ToStringArray(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	file.Close()

	return lines
}