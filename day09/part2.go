package main

import (
	"fmt"
	"strings"
)

func Part2(raw string) {
	var solve func(raw string) int
	solve = func(raw string) int {
		i := strings.IndexByte(raw, '(')
		if i == -1 {
			return len(raw)
		}
		j := strings.IndexByte(raw[i:], ')') + i
		var l, n int
		fmt.Sscanf(raw[i:j], "(%dx%d", &l, &n)
		return i + n*solve(raw[j+1:j+1+l]) + solve(raw[j+1+l:])
	}
	fmt.Println(solve(raw))
}
