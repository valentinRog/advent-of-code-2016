package main

import (
	"fmt"
	"strings"
	"strconv"
)

func Part2(raw string) {
	res := 0
	raw = strings.ReplaceAll(raw, "\n", " ")
	arr := strings.Fields(raw)
	for i := 0; i < len(arr); i += 9 {
		for j := 0; j < 3; j++ {
			a, _ := strconv.Atoi(arr[i+j])
			b, _ := strconv.Atoi(arr[i+j+3])
			c, _ := strconv.Atoi(arr[i+j+6])
			if a+b > c && a+c > b && b+c > a {
				res++
			}
		}
	}
	fmt.Println(res)
}
