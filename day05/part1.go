package main

import (
	"crypto/md5"
	"fmt"
)

func Part1(raw string) {
	pass := ""
	n := 0
	for i := 0; i < 8; i++ {
		for ; ; n++ {
			h := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s%d", raw, n))))
			if h[:5] == "00000" {
				pass += string(h[5])
				n++
				break
			}
		}
	}
	fmt.Println(pass)
}
