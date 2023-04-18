package main

import (
	"fmt"
	"sort"
	"strings"
)

func Part2(raw string) {
	bots := map[int][]int{}
	lowToOut := map[int]int{}
	lowToBot := map[int]int{}
	highToBot := map[int]int{}
	highToOut := map[int]int{}
	for _, line := range strings.Split(raw, "\n") {
		if strings.HasPrefix(line, "value") {
			var k, v int
			fmt.Sscanf(line, "value %d goes to bot %d", &v, &k)
			if _, ok := bots[k]; !ok {
				bots[k] = []int{}
			}
			bots[k] = append(bots[k], v)
		} else {
			var bot int
			fmt.Sscanf(line, "bot %d", &bot)
			var i, n int
			if i = strings.Index(line, "low to output"); i != -1 {
				fmt.Sscanf(line[i:], "low to output %d", &n)
				lowToOut[bot] = n
			} else {
				i = strings.Index(line, "low to bot")
				fmt.Sscanf(line[i:], "low to bot %d", &n)
				lowToBot[bot] = n
			}
			if i = strings.Index(line, "high to bot"); i != -1 {
				fmt.Sscanf(line[i:], "high to bot %d", &n)
				highToBot[bot] = n
			} else {
				i = strings.Index(line, "high to output")
				fmt.Sscanf(line[i:], "high to output %d", &n)
				highToOut[bot] = n
			}
		}
	}
	res := 1
	var solve func(bot int)
	solve = func(bot int) {
		if len(bots[bot]) != 2 {
			return
		}
		sort.Ints(bots[bot])
		low, high := bots[bot][0], bots[bot][1]
		bots[bot] = bots[bot][:0]
		if nb, ok := lowToBot[bot]; ok {
			if _, ok := bots[nb]; !ok {
				bots[nb] = []int{}
			}
			bots[nb] = append(bots[nb], low)
			solve(nb)
		} else if out := lowToOut[bot]; out >= 0 && out <= 2 {
			res *= low
		}
		if nb, ok := highToBot[bot]; ok {
			if _, ok := bots[nb]; !ok {
				bots[nb] = []int{}
			}
			bots[nb] = append(bots[nb], high)
			solve(nb)
		} else if out := highToOut[bot]; out >= 0 && out <= 2 {
			res *= high
		}
	}
	for k, v := range bots {
		if len(v) == 2 {
			solve(k)
			break
		}
	}
	fmt.Println(res)
}
