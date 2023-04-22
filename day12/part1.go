package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(raw string) {
	registers := map[string]int{
		"a": 7,
		"b": 0,
		"c": 0,
		"d": 0,
	}
	arrs := [][]string{}
	for _, line := range strings.Split(raw, "\n") {
		arrs = append(arrs, strings.Fields(line))
	}
	for i := 0; i < len(arrs); i++ {
		arr := arrs[i]
		switch arr[0] {
		case "cpy":
			if v, ok := registers[arr[1]]; ok {
				registers[arr[2]] = v
			} else {
				registers[arr[2]], _ = strconv.Atoi(arr[1])
			}
		case "inc":
			registers[arr[1]]++
		case "dec":
			registers[arr[1]]--
		case "jnz":
			v, ok := registers[arr[1]]
			if !ok {
				v, _ = strconv.Atoi(arr[1])
			}
			if v != 0 {
				v, ok = registers[arr[2]]
				if !ok {
					v, _ = strconv.Atoi(arr[2])
				}
				i += v - 1
			}
		}
	}
	fmt.Println(registers["a"])
}
