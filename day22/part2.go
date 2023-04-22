package main

import (
	"container/list"
	"fmt"
	"math"
	"strings"
)

func Part2(raw string) {
	var zMax complex128
	var hole complex128
	walls := map[complex128]struct{}{}
	{
		arr := strings.Split(raw, "\n")[2:]
		for i, line := range strings.Split(raw, "\n")[2:] {
			var x, y, size int
			fmt.Sscanf(line, "/dev/grid/node-x%d-y%d %dT", &x, &y, &size)
			if size > 200 {
				walls[complex(float64(x), float64(y))] = struct{}{}
			}
			if i == len(arr)-1 {
				zMax = complex(float64(x), float64(y))
			}
			if strings.HasSuffix(line, " 0%") {
				hole = complex(float64(x), float64(y))
			}
		}
	}

	steps := 0
	{
		type State struct {
			hole  complex128
			steps int
		}
		visited := map[complex128]struct{}{
			hole: {},
		}
		q := list.New()
		q.PushBack(State{hole, 0})
		for {
			state := q.Remove(q.Front()).(State)
			if math.Abs(real(state.hole)-real(zMax)) <= 1 && imag(state.hole) <= 1 {
				steps, hole = state.steps, state.hole
				break
			}
			for _, d := range []complex128{1, -1, 1i, -1i} {
				nh := state.hole + d
				if _, ok := walls[nh]; ok {
					continue
				}
				if _, ok := visited[nh]; ok {
					continue
				}
				if real(nh) < 0 || imag(nh) < 0 || real(nh) > real(zMax) || imag(nh) > imag(zMax) {
					continue
				}
				visited[nh] = struct{}{}
				q.PushBack(State{nh, state.steps + 1})
			}
		}
	}
	{
		type State struct {
			data           complex128
			hole           complex128
			steps          int
			last_data_move int
		}
		visited := map[complex128]struct{}{
			complex(real(zMax), 0): {},
		}
		q := list.New()
		q.PushBack(State{complex(real(zMax), 0), hole, 0, 0})
		for {
			state := q.Remove(q.Front()).(State)
			if state.data == complex(0, 0) {
				steps += state.steps
				break
			}
			if state.last_data_move > 4 {
				continue
			}
			for _, d := range []complex128{1, -1, 1i, -1i} {
				nh := state.hole + d
				if _, ok := walls[nh]; ok {
					continue
				}
				if real(nh) < 0 || imag(nh) < 0 || real(nh) > real(zMax) || imag(nh) > imag(zMax) {
					continue
				}
				if nh == state.data {
					nData := state.hole
					if _, ok := visited[nData]; !ok {
						visited[nData] = struct{}{}
						q.PushBack(State{nData, nh, state.steps + 1, 0})
					}
				} else {
					q.PushBack(State{state.data, nh, state.steps + 1, state.last_data_move + 1})
				}
			}
		}
	}
	fmt.Println(steps)
}
