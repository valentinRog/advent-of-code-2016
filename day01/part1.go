package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(raw string) {
	arr := strings.Split(raw, ", ")
	rotations := map[byte]complex128{
		'L': complex(0, 1),
		'R': complex(0, -1),
	}
	d := complex(1, 0)
	p := complex(0, 0)
	for _, s := range arr {
		d *= rotations[s[0]]
		l, _ := strconv.Atoi(s[1:])
		p += d * complex(float64(l), 0)
	}
	fmt.Println(int(math.Abs(real(p)) + math.Abs(imag(p))))
}
