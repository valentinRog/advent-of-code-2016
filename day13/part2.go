package main

import (
	"container/list"
	"fmt"
	"math/bits"
)

func Part2(raw string) {
	var n uint
	fmt.Sscanf(raw, "%d", &n)

	isWall := func(z complex64) bool {
		x, y := uint(real(z)), uint(imag(z))
		return bits.OnesCount(x*x+3*x+2*x*y+y+y*y+uint(n))%2 == 1
	}

	visited := map[complex64]struct{}{
		1 + 1i: {},
	}

	type Pos struct {
		z complex64
		d int
	}

	q := list.New()
	q.PushBack(Pos{1 + 1i, 0})

	for {
		p := q.Remove(q.Front()).(Pos)
		if p.d == 50 {
			fmt.Println(len(visited))
			break
		}
		for _, dir := range []complex64{1, -1, 1i, -1i} {
			z := p.z + dir
			if _, ok := visited[z]; !ok && !isWall(z) && real(z) >= 0 && imag(z) >= 0 {
				visited[z] = struct{}{}
				q.PushBack(Pos{z, p.d + 1})
			}
		}
	}
}
