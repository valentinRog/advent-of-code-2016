package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	res := 0
	for i := 0; i < len(raw); i++ {
		if raw[i] == '(' {
			var l, n int
			fmt.Sscanf(raw[i:], "(%dx%d", &l, &n)
			i += strings.IndexByte(raw[i:], ')') + 1
			res += l * n
			i += l - 1
		} else {
			res++
		}
	}
	fmt.Println(res)
}
