package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func Part2(raw string) {
	n, _ := strconv.Atoi(raw)
	l1 := list.New()
	l2 := list.New()
	for i := 0; i < n; i++ {
		if i < n/2 {
			l1.PushBack(i)
		} else {
			l2.PushBack(i)
		}
	}
	for l1.Len()+l2.Len() > 1 {
		l2.Remove(l2.Front())
		l2.PushBack(l1.Remove(l1.Front()).(int))
		if l1.Len()+1 < l2.Len() {
			l1.PushBack(l2.Remove(l2.Front()).(int))
		}
	}
	fmt.Println(l2.Front().Value.(int) + 1)
}
