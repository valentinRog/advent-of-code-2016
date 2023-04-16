package main

import (
	"crypto/md5"
	"fmt"
)

func Part2(raw string) {
	m := make(map[byte]byte)
	n := 0
	for i := 0; i < 8; i++ {
		for ; ; n++ {
			h := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", raw, n))))
			if _, ok := m[h[5]]; !ok && h[:5] == "00000" && h[5] >= '0' && h[5] <= '7' {
				m[h[5]] = h[6]
				n++
				break
			}
		}
	}
	pass := make([]byte, 8)
	for k, v := range m {
		pass[k-'0'] = v
	}
	fmt.Println(string(pass))
}
