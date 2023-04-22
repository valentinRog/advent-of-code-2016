package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	arr := strings.Split(raw, "\n")[2:]

	type Node struct {
		used      int
		available int
	}
	parseNode := func(s string) Node {
		s = strings.Join(strings.Fields(s)[2:4], "")
		var n Node
		fmt.Sscanf(s, "%dT%dT", &n.used, &n.available)
		return n
	}

	res := 0
	for i, s1 := range arr {
		a := parseNode(s1)
		for j, s2 := range arr {
			if i != j {
				b := parseNode(s2)
				if a.used > 0 && a.used <= b.available {
					res++
				}
			}
		}
	}
	fmt.Println(res)
}
