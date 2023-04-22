package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	ops := map[rune]complex128{
		'U': -1i,
		'D': 1i,
		'L': -1,
		'R': 1,
	}
	p := 1 + 1i
	W := 3.0
	for _, line := range strings.Split(raw, "\n") {
		for _, c := range line {
			np := p + ops[c]
			if real(np) >= 0 && real(np) < W && imag(np) >= 0 && imag(np) < W {
				p = np
			}
		}
		fmt.Print(real(p) + imag(p)*W + 1)
	}
	fmt.Println()
}
