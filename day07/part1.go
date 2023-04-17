package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	res := 0
	for _, line := range strings.Split(raw, "\n") {
		var ins []string
		var outs []string
		sb := strings.Builder{}
		for _, c := range line {
			if c == '[' {
				outs = append(outs, sb.String())
				sb.Reset()
			} else if c == ']' {
				ins = append(ins, sb.String())
				sb.Reset()
			} else {
				sb.WriteRune(c)
			}
		}
		outs = append(outs, sb.String())

		has_abba := func(ss []string) bool {
			for _, s := range ss {
				for i := 0; i < len(s)-3; i++ {
					if s[i] == s[i+3] && s[i+1] == s[i+2] && s[i] != s[i+1] {
						return true
					}
				}
			}
			return false
		}

		if has_abba(outs) && !has_abba(ins) {
			res++
		}
	}
	fmt.Println(res)
}
