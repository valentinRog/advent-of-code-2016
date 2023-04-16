package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part2(raw string) {
	arr := strings.Split(raw, ", ")
	rotations := map[byte]complex128{
		'L': complex(0, 1),
		'R': complex(0, -1),
	}
	d := complex(1, 0)
	p := complex(0, 0)
	visited := map[complex128]struct{}{
		p: {},
	}
out:
	for _, s := range arr {
		d *= rotations[s[0]]
		l, _ := strconv.Atoi(s[1:])
		for i := 0; i < l; i++ {
			p += d
			if _, ok := visited[p]; ok {
				break out
			}
			visited[p] = struct{}{}
		}
	}
	fmt.Println(int(math.Abs(real(p)) + math.Abs(imag(p))))
}
