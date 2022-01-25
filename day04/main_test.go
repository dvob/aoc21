package main

import (
	"bytes"
	"testing"
)

var exampleInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func Test_part1(t *testing.T) {
	b := bytes.NewBufferString(exampleInput)

	boards, numbers, err := read(b)
	if err != nil {
		t.Fatal(err)
	}
	if len(boards) != 3 {
		t.Fatal("boards != 3")
	}

	noWinner := numbers[:11]
	winners := findWinner(noWinner, boards)
	if len(winners) != 0 {
		t.Fatal("expected zero winners")
	}

	score, number, result := findWinnerAll(numbers, boards)

	if number != 24 {
		t.Fatalf("expected 24 as winning number. got=%d", number)
	}
	if score != 188 {
		t.Fatalf("expected score to be 188. got=%d", score)
	}

	if result != 4512 {
		t.Fatalf("expected result to be 4512. got=%d", result)
	}

}

func Test_part2(t *testing.T) {
	b := bytes.NewBufferString(exampleInput)

	boards, numbers, err := read(b)
	if err != nil {
		t.Fatal(err)
	}

	score, number, result := findLastWinner(numbers, boards)

	if number != 13 {
		t.Fatalf("expected 13 as winning number. got=%d", number)
	}
	if score != 148 {
		t.Fatalf("expected score to be 148. got=%d", score)
	}

	if result != 1924 {
		t.Fatalf("expected result to be 1924. got=%d", result)
	}
}
