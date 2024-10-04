package main

import (
	"fmt"
	"sort"
)

func main() {
	// 创建一个int64类型的切片
	numbers := []int64{5, 1, 3, 2, 4}

	// 对切片进行排序
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j] // 升序排序
	})

	// 输出排序后的切片
	fmt.Println(numbers) // 输出: [1 2 3 4 5]
}
