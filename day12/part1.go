package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(raw string) {
	registers := map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
	}

	instructions := map[string]func(arr []string) int{
		"cpy": func(arr []string) int {
			if v, ok := registers[arr[0]]; ok {
				registers[arr[1]] = v
			} else {
				registers[arr[1]], _ = strconv.Atoi(arr[0])
			}
			return 1
		},
		"inc": func(arr []string) int {
			registers[arr[0]]++
			return 1
		},
		"dec": func(arr []string) int {
			registers[arr[0]]--
			return 1
		},
		"jnz": func(arr []string) int {
			v, ok := registers[arr[0]]
			if !ok {
				v, _ = strconv.Atoi(arr[0])
			}
			if v != 0 {
				j, _ := strconv.Atoi(arr[1])
				return j
			}
			return 1
		},
	}

	lines := strings.Split(raw, "\n")
	for i := 0; i < len(lines); {
		arr := strings.Fields(lines[i])
		i += instructions[arr[0]](arr[1:])
	}
	fmt.Println(registers["a"])
}
