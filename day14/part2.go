package main

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Part2(raw string) {
	hashes := map[int]string{}
	hash := func(i int) string {
		if h, ok := hashes[i]; ok {
			return h
		}
		h := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprint(raw, i))))
		for i := 0; i < 2016; i++ {
			h = fmt.Sprintf("%x", md5.Sum([]byte(h)))
		}
		hashes[i] = h
		return h
	}

	isKey := func(i int) bool {
		h := hash(i)
		for j := 0; j < len(h)-2; j++ {
			if h[j] == h[j+1] && h[j] == h[j+2] {
				for k := 1; k <= 1000; k++ {
					nh := hash(i + k)
					if strings.Contains(nh, strings.Repeat(string(h[j]), 5)) {
						return true
					}
				}
				break
			}
		}
		return false
	}

	i := 0
	for n := 0; n < 64; i++ {
		if isKey(i) {
			n++
		}
	}
	fmt.Println(i - 1)
}
