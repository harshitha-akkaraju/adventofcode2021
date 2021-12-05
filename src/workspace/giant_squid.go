package workspace

import (
	"log"
	"math"
	"strconv"
	"strings"
)

/// Day 4 :: Giant Squid
// https://adventofcode.com/2021/day/4

func markBoard(board [][]int, pick int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == pick {
				board[i][j] = -1
				return
			}
		}
	}
}

func findScore(picks []int, board [][]int) (int, int) {
	pickIndex := 0
	pick := 0
	for !isWinning(board) && pickIndex < len(picks) {
		pick = picks[pickIndex]
		markBoard(board, pick)
		pickIndex++
	}

	sum := 0
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != -1 {
				sum += board[i][j]
			}
		}
	}

	return pickIndex, sum * pick
}

func isWinning(board [][]int) bool {
	// check rows
	for _, row := range board {
		if row[0] == -1 &&
			row[1] == -1 &&
			row[2] == -1 &&
			row[3] == -1 &&
			row[4] == -1 {
			return true
		}
	}

	// check cols
	for i := 0; i < len(board[0]); i++ {
		if board[0][i] == -1 &&
			board[1][i] == -1 &&
			board[2][i] == -1 &&
			board[3][i] == -1 &&
			board[4][i] == -1 {
			return true
		}
	}

	return false
}

func findFastestWinningScore(input []string) int {
	picks := []int{}
	board := [][]int{}
	minPickNumber := math.MaxInt
	winningScore := 0

	for _, line := range input {
		if strings.Contains(line, ",") {
			picks = strToIntArr(strings.Split(line, ","))
			continue
		} else if len(line) == 0 && len(board) > 0 {
			pickNumber, score := findScore(picks, board)
			if pickNumber < minPickNumber {
				minPickNumber = pickNumber
				winningScore = score
			}

			board = [][]int{}
			continue
		} else if len(line) > 0 {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "  ", " ")

			elems := strToIntArr(strings.Split(line, " "))
			board = append(board, elems)
		}
	}

	pickNumber, score := findScore(picks, board)
	if pickNumber < minPickNumber {
		minPickNumber = pickNumber
		winningScore = score
	}

	return winningScore
}

func findSlowestWinningScore(input []string) int {
	picks := []int{}
	board := [][]int{}
	maxPickNumber := math.MinInt
	winningScore := 0

	for _, line := range input {
		if strings.Contains(line, ",") {
			picks = strToIntArr(strings.Split(line, ","))
			continue
		} else if len(line) == 0 && len(board) > 0 {
			pickNumber, score := findScore(picks, board)
			if pickNumber > maxPickNumber {
				maxPickNumber = pickNumber
				winningScore = score
			}

			board = [][]int{}
			continue
		} else if len(line) > 0 {
			line = strings.TrimSpace(line)
			line = strings.ReplaceAll(line, "  ", " ")

			elems := strToIntArr(strings.Split(line, " "))
			board = append(board, elems)
		}
	}

	pickNumber, score := findScore(picks, board)
	if pickNumber > maxPickNumber {
		maxPickNumber = pickNumber
		winningScore = score
	}

	return winningScore
}

func strToIntArr(strArr []string) []int {
	elems := []int{}

	for _, elem := range strArr {
		n, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatalln(err.Error())
		}
		elems = append(elems, n)
	}

	return elems
}
