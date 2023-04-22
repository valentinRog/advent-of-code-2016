package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part2(raw string) {
	registers := map[string]int{
		"a": 12,
		"b": 0,
		"c": 0,
		"d": 0,
	}
	lines := strings.Split(raw, "\n")
	for i := 0; i < len(lines); i++ {
		arr := strings.Fields(lines[i])
		switch arr[0] {
		case "cpy":
			if _, ok := registers[arr[2]]; !ok {
				break
			}
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
		case "tgl":
			v, ok := registers[arr[1]]
			if !ok {
				v, _ = strconv.Atoi(arr[1])
			}
			if i+v < len(lines) {
				arr2 := strings.Fields(lines[i+v])
				switch arr2[0] {
				case "inc":
					arr2[0] = "dec"
				case "dec", "tgl":
					arr2[0] = "inc"
				case "jnz":
					arr2[0] = "cpy"
				case "cpy":
					arr2[0] = "jnz"
				}
				lines[i+v] = strings.Join(arr2, " ")
			}
		}
	}
	fmt.Println(registers["a"])
}
