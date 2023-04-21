package main

import (
	"fmt"
	"sort"
	"strings"
)

func Part1(raw string) {
	arr := [][2]int{}
	for _, line := range strings.Split(raw, "\n") {
		var a, b int
		fmt.Sscanf(line, "%d-%d", &a, &b)
		arr = append(arr, [2]int{a, b})
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})

	max := 0
	for _, v := range arr {
		if v[0] > max+1 {
			fmt.Println(max + 1)
			break
		}
		if v[1] > max {
			max = v[1]
		}
	}
}
