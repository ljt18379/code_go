package main

import (
	"sort"
)

func MapSort(m map[string]int, n int) map[string]int {
	value := make([]int, len(m))
	i := 0
	for _, val := range m {
		value[i] = val
		i++
	}
	sort.Ints(value)                              //默认升序
	sort.Sort(sort.Reverse(sort.IntSlice(value))) //降序排序
	mm := make(map[string]int)
	for j := 0; j < n; j++ {
		for k, vv := range m {
			if vv == value[j] {
				mm[k] = vv
				delete(m, k)
				break
			}
		}
	}
	return mm
}
