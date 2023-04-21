package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func Part1(raw string) {
	n, _ := strconv.Atoi(raw)
	r := ring.New(n)
	for i := 0; i < n; i++ {
		r.Value = i
		r = r.Next()
	}
	for r.Next() != r {
		r.Unlink(1)
		r = r.Next()
	}
	fmt.Println(r.Value.(int) + 1)
}
