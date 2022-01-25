package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	notes, err := readInput(f)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("part one:", solvePartOne(notes))

	sum := 0
	for _, note := range notes {
		sum += note.solve()
	}
	fmt.Println("part two:", sum)
}

type note struct {
	inputs  [][]byte
	outputs [][]byte
}

func overlap(a, b []byte) int {
	n := 0
	for _, elementA := range a {
		for _, elementB := range b {
			if elementA == elementB {
				n++
			}
		}
	}
	return n
}

func (n *note) solve() int {
	digits := make([]int, 0, 4)
	for _, output := range n.outputs {
		digits = append(digits, n.solveOutput(output))
	}
	return intsToInt(digits)
}

func (n *note) solveOutput(output []byte) int {
	one := n.getByLen(2)
	four := n.getByLen(4)
	seven := n.getByLen(3)
	if one == nil || four == nil || seven == nil {
		panic("got not one number of each")
	}
	switch len(output) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5:
		if overlap(output, one) == 2 {
			return 3
		}
		if overlap(output, four) == 3 {
			return 5
		}
		return 2
	case 6:
		if overlap(output, seven) == 2 {
			return 6
		}
		if overlap(output, four) == 4 {
			return 9
		}
		return 0
	case 7:
		return 8
	}
	fmt.Println(string(output), n.String())
	panic("failed to solve")
}

func (n *note) getByLen(l int) []byte {
	for _, input := range n.inputs {
		if len(input) == l {
			return input
		}
	}
	return nil
}

func (n *note) String() string {
	buf := &bytes.Buffer{}
	buf.Write(bytes.Join(n.inputs, []byte{' '}))
	buf.WriteString(" | ")
	buf.Write(bytes.Join(n.outputs, []byte{' '}))
	return buf.String()
}

func solve(note note) int {
	return note.solve()
}

func intsToInt(ints []int) int {
	r := 0
	for i := 1; i <= len(ints); i++ {
		m := 1
		for j := 1; j < i; j++ {
			m *= 10
		}
		r += m * ints[len(ints)-i]
	}
	return r
}

func solvePartOne(notes []note) int {
	count := 0
	for _, note := range notes {
		for _, output := range note.outputs {
			l := len(output)
			if l == 2 || l == 4 || l == 3 || l == 7 {
				count++
			}
		}
	}
	return count
}

func readInput(input io.Reader) ([]note, error) {
	s := bufio.NewScanner(input)

	notes := []note{}
	for s.Scan() {
		data := make([]byte, len(s.Bytes()))
		copy(data, s.Bytes())
		parts := bytes.SplitN(data, []byte{'|'}, 2)
		inputs := bytes.Fields(parts[0])
		outputs := bytes.Fields(parts[1])
		notes = append(notes, note{inputs, outputs})
	}
	if s.Err() != nil {
		return nil, s.Err()
	}
	return notes, nil
}
