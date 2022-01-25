package main

import (
	"bytes"
	"testing"
)

func Test_solve(t *testing.T) {
	input := bytes.NewBufferString("acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")

	notes, err := readInput(input)
	if err != nil {
		t.Fatal(err)
	}

	//t.Log(notes[0].String())

	expected := 5353
	result := solve(notes[0])
	if result != expected {
		t.Fatalf("expected=%d, got=%d", expected, result)
	}
}

func Test_intsToInt(t *testing.T) {
	input := []int{5, 3, 5, 3}
	result := intsToInt(input)
	if result != 5353 {
		t.Fatalf("got=%d, want=%d", result, 5353)
	}
}

func Test_overlap(t *testing.T) {
	a := []byte{'a', 'b', 'c'}
	b := []byte{'c', 'x', 'b'}
	result := overlap(a, b)
	if result != 2 {
		t.Fatalf("got=%d, want=%d", result, 2)
	}
}
