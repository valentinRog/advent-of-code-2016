package main

import (
	"fmt"
	"strings"
)

func Part2(raw string) {
	const W = 50
	const H = 6
	m := map[complex64]struct{}{}
	for _, line := range strings.Split(raw, "\n") {
		if strings.HasPrefix(line, "rect") {
			var w, h int
			fmt.Sscanf(line, "rect %dx%d", &w, &h)
			for x := 0; x < w; x++ {
				for y := 0; y < h; y++ {
					m[complex(float32(x), float32(y))] = struct{}{}
				}
			}
		} else if strings.HasPrefix(line, "rotate row") {
			var y, n int
			fmt.Sscanf(line, "rotate row y=%d by %d", &y, &n)
			var tmp []complex64
			for x := 0; x < W; x++ {
				z := complex(float32(x), float32(y))
				if _, ok := m[z]; ok {
					tmp = append(tmp, z)
					delete(m, z)
				}
			}
			for _, z := range tmp {
				r := float32((int(real(z)) + n) % W)
				m[complex(r, float32(y))] = struct{}{}
			}
		} else {
			var x, n int
			fmt.Sscanf(line, "rotate column x=%d by %d", &x, &n)
			var tmp []complex64
			for y := 0; y < H; y++ {
				z := complex(float32(x), float32(y))
				if _, ok := m[z]; ok {
					tmp = append(tmp, z)
					delete(m, z)
				}
			}
			for _, z := range tmp {
				r := float32((int(imag(z)) + n) % H)
				m[complex(float32(x), r)] = struct{}{}
			}
		}
	}
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			z := complex(float32(x), float32(y))
			if _, ok := m[z]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
