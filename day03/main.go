package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func readBin(text string) (uint64, error) {
	if len(text) > 64 {
		return 0, fmt.Errorf("cant read binary '%s': too long", text)
	}
	var r uint64 = 0
	for i := len(text) - 1; i >= 0; i-- {
		switch text[i] {
		case '1':
			r = r | 1<<(len(text)-(i+1))
		case '0':
		default:
			return 0, fmt.Errorf("unrecognized character: %c", text[i])
		}
	}
	return r, nil
}

func read(input io.Reader) ([]uint64, int, error) {
	scanner := bufio.NewScanner(input)
	numbers := []uint64{}
	first := true
	l := 0
	for scanner.Scan() {
		text := scanner.Text()
		if first {
			l = len(text)
		}
		if len(text) != l {
			return nil, 0, fmt.Errorf("line with different length")
		}
		val, err := readBin(text)
		if err != nil {
			return nil, 0, err
		}
		numbers = append(numbers, val)
	}
	if scanner.Err() != nil {
		return nil, 0, scanner.Err()
	}
	return numbers, l, nil
}

func solvePartOne(numbers []uint64, length int) (gamma uint64, epsilon uint64, result uint64) {
	positions := make([]int, length)
	for _, n := range numbers {
		for i := 0; i < length; i++ {
			if isSet(n, i+1) {
				positions[i] += 1
			} else {
				positions[i] -= 1
			}
		}
	}
	for i, v := range positions {
		if v > 0 {
			gamma = gamma | 1<<(i)
		} else {
			epsilon = epsilon | 1<<(i)
		}
	}
	return gamma, epsilon, gamma * epsilon
}

func filterPos(numbers []uint64, pos int, ones bool) []uint64 {
	newNumbers := []uint64{}
	for _, n := range numbers {
		if ones == isSet(n, pos) {
			newNumbers = append(newNumbers, n)
		}
	}
	return newNumbers
}

func isSet(n uint64, pos int) bool {
	if pos > 64 {
		panic("position to high")
	}
	return (n & (1 << (pos - 1))) != 0
}

func countOnes(numbers []uint64, pos int) int {
	i := 0
	for _, n := range numbers {
		if isSet(n, pos) {
			i++
		}
	}
	return i
}

func filter(numbers []uint64, l int, oxygen bool) uint64 {
	for i := l; i > 0; i-- {
		ones := countOnes(numbers, i)
		zeros := len(numbers) - ones
		if oxygen {
			if ones >= zeros {
				numbers = filterPos(numbers, i, true)
			} else {
				numbers = filterPos(numbers, i, false)
			}
		} else {
			if ones < zeros {
				numbers = filterPos(numbers, i, true)
			} else {
				numbers = filterPos(numbers, i, false)
			}
		}
		if len(numbers) == 1 {
			break
		}
	}
	return numbers[0]
}

func solvePartTwo(numbers []uint64, length int) (oxygen uint64, co2 uint64, result uint64) {
	oxygen = filter(numbers, length, true)
	co2 = filter(numbers, length, false)
	return oxygen, co2, oxygen * co2
}

func main() {
	input, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	numbers, l, err := read(input)
	if err != nil {
		log.Fatal(err)
	}

	gamma, epsilon, result := solvePartOne(numbers, l)
	fmt.Printf("gamma %d %#b\n", gamma, gamma)
	fmt.Printf("epsilon %d %#b\n", epsilon, epsilon)
	fmt.Printf("result part1: %d\n", result)

	oxygen, co2, result2 := solvePartTwo(numbers, l)
	fmt.Printf("oxygen %d %#b\n", oxygen, oxygen)
	fmt.Printf("co2 %d %#b\n", co2, co2)
	fmt.Printf("result part2: %d\n", result2)
}
