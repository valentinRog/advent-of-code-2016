package main

import (
	"fmt"
	"strings"
)

func Part2(raw string) {
	traps := map[string]struct{}{
		"^^.": {},
		".^^": {},
		"^..": {},
		"..^": {},
	}

	getByte := func(s string) byte {
		if _, ok := traps[s]; ok {
			return '^'
		}
		return '.'
	}

	res := strings.Count(raw, ".")
	for i := 0; i < 400000-1; i++ {
		var sb strings.Builder
		sb.WriteByte(getByte("." + raw[:2]))
		for i := 1; i < len(raw)-1; i++ {
			sb.WriteByte(getByte(raw[i-1 : i+2]))
		}
		sb.WriteByte(getByte(raw[len(raw)-2:] + "."))
		raw = sb.String()
		res += strings.Count(raw, ".")
	}
	fmt.Println(res)
}
