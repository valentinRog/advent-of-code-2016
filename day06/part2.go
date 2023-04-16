package main

import (
	"fmt"
	"strings"
)

func Part2(raw string) {
	arr := strings.Split(raw, "\n")
	for i := 0; i < len(arr[0]); i++ {
		counter := [26]int{}
		for _, line := range arr {
			counter[line[i]-'a']++
		}
		var c byte
		min := len(arr)
		for j, n := range counter {
			if n < min {
				min = n
				c = byte(j) + 'a'
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Println()
}
