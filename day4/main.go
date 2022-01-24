package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type board [5][5]uint8

// func (b *board) isLoser(occured []uint8) bool {
// 	return !b.hasRowWinner(occured) && !b.hasColWinner(occured)
// }

func (b *board) hasWinner(occured []uint8) bool {
	return b.hasRowWinner(occured) || b.hasColWinner(occured)
}

func (b *board) hasColWinner(occured []uint8) bool {
	for col := 0; col < 5; col++ {
		colOk := true
		for row := 0; row < 5; row++ {
			rowOk := false
			for _, n := range occured {
				if n == b[row][col] {
					rowOk = true
					break
				}
			}
			if !rowOk {
				colOk = false
				break
			}
		}
		if colOk {
			return true
		}
	}
	return false
}

func (b *board) hasRowWinner(occured []uint8) bool {
	for _, row := range *b {
		rowOk := true
		for _, col := range row {
			colOk := false
			for _, n := range occured {
				if n == col {
					colOk = true
					break
				}
			}
			if !colOk {
				rowOk = false
				break
			}
		}
		if rowOk {
			return true
		}
	}
	return false
}

func (b *board) score(occured []uint8) int {
	score := 0
	for _, row := range *b {
		for _, col := range row {
			found := false
			for _, num := range occured {
				if col == num {
					found = true
					continue
				}
			}
			if !found {
				score += int(col)
			}
		}
	}
	return score
}

func findWinner(numbers []uint8, boards []board) []int {
	indexes := []int{}
	for i, board := range boards {
		if board.hasWinner(numbers) {
			indexes = append(indexes, i)
		}
	}
	return indexes
}

func findWinnerAll(numbers []uint8, boards []board) (score int, winningNumber uint8, result int) {
	for i := 1; i <= len(numbers); i++ {
		winners := findWinner(numbers[:i], boards)
		if len(winners) != 0 {
			score = boards[winners[0]].score(numbers[:i])
			winningNumber = numbers[i-1]
			result = score * int(winningNumber)
			return
		}
	}
	return 0, 0, 0
}

func findLastWinner(numbers []uint8, boards []board) (score int, winningNumber uint8, result int) {
	for i := len(numbers); i >= 1; i-- {
		for _, board := range boards {
			if !board.hasWinner(numbers[:i]) {
				score = board.score(numbers[:i+1])
				winningNumber = numbers[i]
				result = score * int(winningNumber)
				return
			}
		}
	}
	return 0, 0, 0
}

func read(input io.Reader) (boards []board, numbers []uint8, err error) {
	scanner := bufio.NewScanner(input)
	if !scanner.Scan() {
		return nil, nil, fmt.Errorf("missing numbers")
	}
	numberStrings := strings.Split(scanner.Text(), ",")
	numbers = make([]uint8, 0, len(numberStrings))
	for _, s := range numberStrings {
		number, err := strconv.Atoi(s)
		if err != nil {
			return nil, nil, err
		}
		numbers = append(numbers, uint8(number))
	}
	boards = []board{}
	currentBoard := board{}
	rowIndex := 0
	for scanner.Scan() {
		if scanner.Text() == "" {
			rowIndex = 0
			currentBoard = board{}
			continue
		}

		numbers := [5]uint8{}
		numberStrings = strings.Fields(scanner.Text())
		if len(numberStrings) != 5 {
			return nil, nil, fmt.Errorf("invalid number of fields")
		}
		for i, str := range numberStrings {
			number, err := strconv.Atoi(str)
			if err != nil {
				return nil, nil, err
			}
			numbers[i] = uint8(number)
		}
		currentBoard[rowIndex] = numbers
		rowIndex++
		if rowIndex == 5 {
			boards = append(boards, currentBoard)
		}
	}
	if scanner.Err() != nil {
		return nil, nil, err
	}
	return boards, numbers, nil
}

func main() {

	boards, numbers, err := read(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("--- part 1 ---")
	score, number, result := findWinnerAll(numbers, boards)
	fmt.Println("number", number)
	fmt.Println("score", score)
	fmt.Println("result", result)

	fmt.Println("--- part 2 ---")
	score, number, result = findLastWinner(numbers, boards)
	fmt.Println("number", number)
	fmt.Println("score", score)
	fmt.Println("result", result)

}
