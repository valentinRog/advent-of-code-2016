package main

import (
	"container/list"
	"fmt"
	"strings"
)

func Part1(raw string) {
	walls := map[complex128]struct{}{}
	checkpoints := []complex128{}
	var start complex128
	for y, line := range strings.Split(raw, "\n") {
		for x, c := range line {
			switch {
			case c == '#':
				walls[complex(float64(x), float64(y))] = struct{}{}
			case c >= '1' && c <= '9':
				checkpoints = append(checkpoints, complex(float64(x), float64(y)))
			case c == '0':
				start = complex(float64(x), float64(y))
			}
		}
	}

	cache := map[[2]complex128]int{}
	dist := func(a, b complex128) int {
		if d, ok := cache[[2]complex128{a, b}]; ok {
			return d
		}
		if d, ok := cache[[2]complex128{b, a}]; ok {
			return d
		}
		type State struct {
			p complex128
			d int
		}
		q := list.New()
		q.PushBack(State{a, 0})
		visited := map[complex128]struct{}{}
		for {
			s := q.Remove(q.Front()).(State)
			for _, d := range []complex128{1, -1, 1i, -1i} {
				p := s.p + d
				if p == b {
					cache[[2]complex128{a, b}] = s.d + 1
					return s.d + 1
				}
				if _, ok := walls[p]; ok {
					continue
				}
				if _, ok := visited[p]; ok {
					continue
				}
				visited[p] = struct{}{}
				q.PushBack(State{p, s.d + 1})
			}
		}
	}

	var solve func([]complex128, uint16) int
	solve = func(path []complex128, remaining uint16) int {
		if remaining == 0 {
			d := 0
			for i := range path[:len(path)-1] {
				d += dist(path[i], path[i+1])
			}
			return d
		}
		min := 1 << 30
		for i := range checkpoints {
			if remaining&(1<<i) == 0 {
				continue
			}
			d := solve(append(path, checkpoints[i]), remaining-(1<<i))
			if d < min {
				min = d
			}
		}
		return min
	}
	fmt.Println(solve([]complex128{start}, (1<<len(checkpoints))-1))
}
