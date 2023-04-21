package main

import (
	"fmt"
)

func Part1(raw string) {
	a := []byte(raw)
	for len(a) < 272 {
		b := make([]byte, len(a))
		copy(b, a)
		for i := 0; i < len(b)/2; i++ {
			b[i], b[len(b)-1-i] = b[len(b)-1-i], b[i]
		}
		for i := range b {
			switch b[i] {
			case '0':
				b[i] = '1'
			case '1':
				b[i] = '0'
			}
		}
		a = append(a, '0')
		a = append(a, b...)
	}
	a = a[:272]
	for len(a)%2 == 0 {
		b := make([]byte, len(a)/2)
		for i := 0; i < len(b); i++ {
			if a[2*i] == a[2*i+1] {
				b[i] = '1'
			} else {
				b[i] = '0'
			}
		}
		a = b
	}
	fmt.Println(string(a))
}
