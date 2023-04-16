package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func Part1(raw string) {
	arr := strings.Split(raw, "\n")
	res := 0
	for _, line := range arr {
		strs := regexp.MustCompile(`[a-z]+`).FindAllString(line, -1)
		name := strings.Join(strs[:len(strs)-1], " ")
		name = strings.ReplaceAll(name, " ", "")
		checksum := strs[len(strs)-1]
		id, _ := strconv.Atoi(regexp.MustCompile(`[0-9]+`).FindString(line))
		counter := make(map[byte]int)
		for _, c := range []byte(name) {
			counter[c]++
		}
		var arr []byte
		for k := range counter {
			arr = append(arr, k)
		}
		sort.Slice(arr, func(i, j int) bool {
			if counter[arr[i]] == counter[arr[j]] {
				return arr[i] < arr[j]
			}
			return counter[arr[i]] > counter[arr[j]]
		})
		arr = arr[:5]
		if string(arr) == checksum {
			res += id
		}
	}
	fmt.Println(res)
}
