package main

import (
	"container/list"
	"crypto/md5"
	"fmt"
)

func Part2(raw string) {
	type Pos struct {
		z    complex64
		path string
	}

	q := list.New()
	q.PushBack(Pos{0, ""})

	res := 0
	for q.Len() > 0 {
		p := q.Remove(q.Front()).(Pos)
		h := fmt.Sprintf("%x", md5.Sum([]byte(raw+p.path)))
		for i, dir := range []complex64{-1i, 1i, -1, 1} {
			if h[i] >= 'b' && h[i] <= 'f' {
				z := p.z + dir
				if real(z) >= 0 && real(z) < 4 && imag(z) >= 0 && imag(z) < 4 {
					np := Pos{z, p.path + string("UDLR"[i])}
					if z == 3+3i {
						if len(np.path) > res {
							res = len(np.path)
						}
					} else {
						q.PushBack(np)
					}
				}
			}
		}
	}
	fmt.Println(res)
}
