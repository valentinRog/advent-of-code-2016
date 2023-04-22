package main

import (
	"fmt"
	"strings"
)

func Part2(raw string) {
	pad := map[complex128]string{
		2 + 0i: "1",
		1 + 1i: "2",
		2 + 1i: "3",
		3 + 1i: "4",
		0 + 2i: "5",
		1 + 2i: "6",
		2 + 2i: "7",
		3 + 2i: "8",
		4 + 2i: "9",
		1 + 3i: "A",
		2 + 3i: "B",
		3 + 3i: "C",
		2 + 4i: "D",
	}
	ops := map[rune]complex128{
		'U': -1i,
		'D': 1i,
		'L': -1,
		'R': 1,
	}
	p := 1 + 1i
	for _, line := range strings.Split(raw, "\n") {
		for _, c := range line {
			np := p + ops[c]
			if _, ok := pad[np]; ok {
				p = np
			}
		}
		fmt.Print(pad[p])
	}
	fmt.Println()
}
