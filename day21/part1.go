package main

import (
	"fmt"
	"strings"
)

func Part1(raw string) {
	s := "abcdefgh"
	ops := map[string]func(string) string{
		"swap position": func(raw string) string {
			var x, y int
			fmt.Sscanf(raw, "swap position %d with position %d", &x, &y)
			b := []byte(s)
			b[x], b[y] = b[y], b[x]
			return string(b)
		},
		"swap letter": func(raw string) string {
			var x, y byte
			fmt.Sscanf(raw, "swap letter %c with letter %c", &x, &y)
			sb := strings.Builder{}
			for _, c := range s {
				switch c {
				case rune(x):
					sb.WriteRune(rune(y))
				case rune(y):
					sb.WriteRune(rune(x))
				default:
					sb.WriteRune(c)
				}
			}
			return sb.String()
		},
		"rotate left": func(raw string) string {
			var x int
			fmt.Sscanf(raw, "rotate left %d steps", &x)
			x %= len(s)
			return s[x:] + s[:x]
		},
		"rotate right": func(raw string) string {
			var x int
			fmt.Sscanf(raw, "rotate right %d steps", &x)
			x %= len(s)
			return s[len(s)-x:] + s[:len(s)-x]
		},
		"rotate based": func(raw string) string {
			var x byte
			fmt.Sscanf(raw, "rotate based on position of letter %c", &x)
			i := strings.IndexByte(s, x) + 1
			if i > 4 {
				i++
			}
			i %= len(s)
			return s[len(s)-i:] + s[:len(s)-i]
		},
		"reverse positions": func(raw string) string {
			var x, y int
			fmt.Sscanf(raw, "reverse positions %d through %d", &x, &y)
			b := []byte(s)
			for i := 0; i < (y-x+1)/2; i++ {
				b[x+i], b[y-i] = b[y-i], b[x+i]
			}
			return string(b)
		},
		"move position": func(raw string) string {
			var x, y int
			fmt.Sscanf(raw, "move position %d to position %d", &x, &y)
			b := []byte(s)
			c := b[x]
			b = append(b[:x], b[x+1:]...)
			b = append(b[:y], append([]byte{c}, b[y:]...)...)
			return string(b)
		},
	}
	for _, line := range strings.Split(raw, "\n") {
		op := strings.Fields(line)[:2]
		s = ops[strings.Join(op, " ")](line)
	}
	fmt.Println(s)
}
