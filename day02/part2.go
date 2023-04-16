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
	update := map[string]func(complex128) complex128{
		"U": func(p complex128) complex128 {
			if _, ok := pad[p-1i]; ok {
				return p - 1i
			}
			return p
		},
		"D": func(p complex128) complex128 {
			if _, ok := pad[p+1i]; ok {
				return p + 1i
			}
			return p
		},
		"L": func(p complex128) complex128 {
			if _, ok := pad[p-1]; ok {
				return p - 1
			}
			return p
		},
		"R": func(p complex128) complex128 {
			if _, ok := pad[p+1]; ok {
				return p + 1
			}
			return p
		},
	}
	p := 2 + 0i
	for _, line := range strings.Split(raw, "\n") {
		for _, c := range line {
			p = update[string(c)](p)
		}
		fmt.Print(pad[p])
	}
	fmt.Println()
}
