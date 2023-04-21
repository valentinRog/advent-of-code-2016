package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	type Disc struct {
		p int
		n int
	}
	discs := []Disc{}
	for _, line := range strings.Split(raw, "\n") {
		line = line[strings.Index(line, "has"):]
		var d Disc
		fmt.Sscanf(line, "has %d positions; at time=0, it is at position %d.", &d.n, &d.p)
		discs = append(discs, d)
	}

out:
	for t := 0; ; t++ {
		for i, d := range discs {
			if (d.p+t+i+1)%d.n != 0 {
				continue out
			}
		}
		fmt.Println(t)
		break
	}
}
