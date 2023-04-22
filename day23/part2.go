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
	arrs := [][]string{}
	for _, line := range strings.Split(raw, "\n") {
		arrs = append(arrs, strings.Fields(line))
	}
	for i := 0; i < len(arrs); i++ {
		arr := arrs[i]
		if arr[0] == "inc" && arr[1] == "a" {
			registers["a"] += registers["d"] * registers["c"]
			registers["d"] = 0
			registers["c"] = 0
			i += 4
			continue
		}
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
			if i+v < len(arrs) {
				switch arrs[i+v][0] {
				case "inc":
					arrs[i+v][0] = "dec"
				case "dec", "tgl":
					arrs[i+v][0] = "inc"
				case "jnz":
					arrs[i+v][0] = "cpy"
				case "cpy":
					arrs[i+v][0] = "jnz"
				}
			}
		}
	}
	fmt.Println(registers["a"])
}
