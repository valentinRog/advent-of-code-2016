package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	res := 0
	for _, line := range strings.Split(raw, "\n") {
		var a, b, c int
		fmt.Sscanf(line, "%d %d %d", &a, &b, &c)
		if a+b > c && a+c > b && b+c > a {
			res++
		}
	}
	fmt.Println(res)
}
