package main

import (
	"fmt"
	"strings"
)

func Part2(raw string) {
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

	next:
		for _, so := range outs {
			for i := 0; i < len(so)-2; i++ {
				if so[i] == so[i+2] && so[i] != so[i+1] {
					for _, si := range ins {
						if strings.Contains(si, string([]byte{so[i+1], so[i], so[i+1]})) {
							res++
							break next
						}
					}
				}
			}
		}
	}
	fmt.Println(res)
}
