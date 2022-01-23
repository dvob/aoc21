package main

import (
	"bytes"
	"testing"
)

func Test_readBin(t *testing.T) {
	tests := []struct {
		input  string
		result uint64
	}{
		{"1001", 9},
	}

	for _, test := range tests {
		v, err := readBin(test.input)
		if err != nil {
			t.Fatal(err)
		}
		if v != test.result {
			t.Fatalf("input: %s, got: %d, want: %d", test.input, v, test.result)
		}
	}
}

func Test_isSet(t *testing.T) {
	// 9 = 1001

	if !isSet(9, 1) {
		t.Fatal("1 should be set in 9")
	}

	if !isSet(9, 4) {
		t.Fatal("4 should be set in 9")
	}
	if isSet(9, 3) {
		t.Fatal("3 should not be set in 9")
	}
}

func Test_isSet1(t *testing.T) {
	i := uint64(3460)

	if !isSet(3460, 3) {
		t.Fatalf("3 should be set in %d", i)
	}
	if isSet(i, 1) {
		t.Fatalf("1 should not be set in %d", i)
	}
}

var exampleInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func Test_part1(t *testing.T) {
	b := bytes.NewBufferString(exampleInput)
	numbers, l, err := read(b)
	if err != nil {
		t.Fatal(err)
	}
	gamma, epsilon, result := solvePartOne(numbers, l)
	if gamma != 22 {
		t.Errorf("gamma: got: %d, expected: %d", gamma, 22)
	}
	if epsilon != 9 {
		t.Errorf("epsilon: got: %d, expected: %d", epsilon, 9)
	}
	if result != 198 {
		t.Errorf("result: got: %d, expected: %d", gamma, 198)
	}
}

func Test_part2(t *testing.T) {
	b := bytes.NewBufferString(exampleInput)
	numbers, l, err := read(b)
	if err != nil {
		t.Fatal(err)
	}
	oxygen, co2, result := solvePartTwo(numbers, l)
	if oxygen != 23 {
		t.Errorf("oxygen: got: %d, expected: %d", oxygen, 23)
	}
	if co2 != 10 {
		t.Errorf("co2: got: %d, expected: %d", co2, 10)
	}
	if result != 230 {
		t.Errorf("result: got: %d, expected: %d", result, 230)
	}
}
