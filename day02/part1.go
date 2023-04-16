package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	W := 3
	update := map[string]func(int) int{
		"U": func(p int) int {
			if p/W > 0 {
				return p - W
			}
			return p
		},
		"D": func(p int) int {
			if p/W < W-1 {
				return p + W
			}
			return p
		},
		"L": func(p int) int {
			if p%W > 0 {
				return p - 1
			}
			return p
		},
		"R": func(p int) int {
			if p%W < W-1 {
				return p + 1
			}
			return p
		},
	}
	p := 1 * W + 1;
	for _, line := range strings.Split(raw, "\n") {
		for _, c := range line {
			p = update[string(c)](p)
		}
		fmt.Print(p + 1)
	}
	fmt.Println()
}
