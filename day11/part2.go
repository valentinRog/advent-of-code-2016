package main

import (
	"container/list"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

func Part2(raw string) {
	generators := map[string]int{}
	microchips := map[string]int{}
	pos := []int{}
	for floor, line := range strings.Split(raw, "\n") {
		arr := strings.Split(line, " ")
		for i, item := range arr[:len(arr)-1] {
			switch {
			case strings.HasPrefix(arr[i+1], "generator"):
				generators[item] = len(pos)
				pos = append(pos, floor)
			case strings.HasPrefix(arr[i+1], "microchip"):
				item = strings.ReplaceAll(item, "-compatible", "")
				microchips[item] = len(pos)
				pos = append(pos, floor)
			}
		}
	}
	generators["elerium"] = len(pos)
	pos = append(pos, 0)
	generators["dilithium"] = len(pos)
	pos = append(pos, 0)
	microchips["elerium"] = len(pos)
	pos = append(pos, 0)
	microchips["dilithium"] = len(pos)
	pos = append(pos, 0)

	valid := func(pos []int) bool {
		for k, v := range microchips {
			shield := false
			for k2, v2 := range generators {
				if k == k2 && pos[v] == pos[v2] {
					shield = true
					break
				}
			}
			if shield {
				continue
			}
			for _, v2 := range generators {
				if pos[v] == pos[v2] {
					return false
				}
			}
		}
		return true
	}

	type state struct {
		pos []int
		e   int
		n   int
	}

	q := list.New()
	q.PushBack(state{pos, 0, 0})

	cache := map[string]struct{}{}
	make_key := func(pos []int, e int) string {
		arr := [][2]int{}
		for k, v := range microchips {
			arr = append(arr, [2]int{pos[v], pos[generators[k]]})
		}
		sort.Slice(arr, func(i, j int) bool {
			if arr[i][0] == arr[j][0] {
				return arr[i][1] < arr[j][1]
			}
			return arr[i][0] < arr[j][0]
		})
		return fmt.Sprint(arr, e)
	}

	target := make([]int, len(pos))
	for i := range target {
		target[i] = 3
	}

out:
	for {
		s := q.Front().Value.(state)
		q.Remove(q.Front())
		for _, e := range []int{s.e - 1, s.e + 1} {
			if e < 0 || e > 3 {
				continue
			}
			for i := range s.pos {
				for j := range s.pos {
					if s.pos[i] != s.e || s.pos[j] != s.e {
						continue
					}
					pos := make([]int, len(s.pos))
					copy(pos, s.pos)
					pos[i] = e
					if i != j {
						pos[j] = e
					}
					if reflect.DeepEqual(pos, target) {
						fmt.Println(s.n + 1)
						break out
					}
					if valid(pos) {
						if _, ok := cache[make_key(pos, e)]; !ok {
							cache[make_key(pos, e)] = struct{}{}
							q.PushBack(state{pos, e, s.n + 1})
						}
					}
				}
			}
		}
	}
}
