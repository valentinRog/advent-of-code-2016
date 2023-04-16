package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part2(raw string) {
	arr := strings.Split(raw, "\n")
	for _, line := range arr {
		strs := regexp.MustCompile(`[a-z]+`).FindAllString(line, -1)
		name := strings.Join(strs[:len(strs)-1], " ")
		id, _ := strconv.Atoi(regexp.MustCompile(`[0-9]+`).FindString(line))
		sb := strings.Builder{}
		for _, c := range name {
			if c != ' ' {
				c = rune((int(c-'a')+id)%26) + 'a'
			}
			sb.WriteRune(c)
		}
		if strings.Contains(sb.String(), "north") {
			fmt.Println(id)
			break
		}
	}
}
