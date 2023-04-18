package main

import (
	"fmt"
	"sort"
	"strings"
)

func Part1(raw string) {
	bots := map[int][]int{}
	lowToOut := map[int]int{}
	lowToBot := map[int]int{}
	highToBot := map[int]int{}
	highToOut := map[int]int{}
	for _, line := range strings.Split(raw, "\n") {
		if strings.HasPrefix(line, "value") {
			var k, v int
			fmt.Sscanf(line, "value %d goes to bot %d", &v, &k)
			if _, ok := bots[k]; ok {
				bots[k] = append(bots[k], v)
			} else {
				bots[k] = []int{v}
			}
		} else {
			var bot int
			fmt.Sscanf(line, "bot %d", &bot)
			var i, n int
			if i = strings.Index(line, "low to output"); i != -1 {
				fmt.Sscanf(line[i:], "low to output %d", &n)
				lowToOut[bot] = n
			}
			if i = strings.Index(line, "low to bot"); i != -1 {
				fmt.Sscanf(line[i:], "low to bot %d", &n)
				lowToBot[bot] = n
			}
			if i = strings.Index(line, "high to bot"); i != -1 {
				fmt.Sscanf(line[i:], "high to bot %d", &n)
				highToBot[bot] = n
			}
			if i = strings.Index(line, "high to output"); i != -1 {
				fmt.Sscanf(line[i:], "high to output %d", &n)
				highToOut[bot] = n
			}
		}
	}
	var solve func(bot int) int
	solve = func(bot int) int {
		return 0
	}
}
