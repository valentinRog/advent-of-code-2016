package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Part1(raw string) {
	compute := func(a int) bool {
		registers := map[string]int{
			"a": a,
			"b": 0,
			"c": 0,
			"d": 0,
		}

		type Config struct {
			a, b, c, d, i, s int
		}
		configs := map[Config]struct{}{}

		signal := 1
		arrs := [][]string{}
		for _, line := range strings.Split(raw, "\n") {
			arrs = append(arrs, strings.Fields(line))
		}
		for i := 0; i < len(arrs); i++ {
			conf := Config{
				a: registers["a"],
				b: registers["b"],
				c: registers["c"],
				d: registers["d"],
				i: i,
				s: signal,
			}
			if _, ok := configs[conf]; ok {
				return true
			}
			configs[conf] = struct{}{}

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
			case "out":
				v, ok := registers[arr[1]]
				if !ok {
					v, _ = strconv.Atoi(arr[1])
				}
				if (v != 0 && v != 1) || v == signal {
					return false
				}
				signal = v
			}
		}
		return false
	}

	i := 0
	for ; !compute(i); i++ {
	}
	fmt.Println(i)
}
