package main

import "reflect"

func KillRepetion(nums [][]interface{}) [][]interface{} {
	newRes := make([][]interface{}, 0)
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if reflect.DeepEqual(nums[i], nums[j]) {
				flag = true
				break
			}
		}
		if !flag {
			newRes = append(newRes, nums[i])
		}
	}
	return newRes
}

func SearchMatrix(matrix [][]interface{}, target []interface{}) (int, bool) {
	if matrix == nil || len(matrix[0]) < 1 {
		return -1, false
	}
	col, row := 0, 0
	//从右上角元素开始，大于目标元素则左移，小于目标元素则下移，直到找到或便利一遍为止
	for row <= len(matrix)-1 && col < len(target) {
		if matrix[row][col] == target[col] {
			col++
			//与SearchMatrix2方法的区别
			if col == len(target)-1 {
				return row, true
			}
		} else {
			col = 0
			row++
			if row >= len(matrix) {
				return -1, false
			}
		}
	}
	return -1, false
}

func SearchMatrix2(matrix [][]interface{}, target []interface{}) (int, bool) {
	if matrix == nil || len(matrix[0]) < 1 {
		return -1, false
	}
	col, row := 0, 0
	//从右上角元素开始，大于目标元素则左移，小于目标元素则下移，直到找到或便利一遍为止
	for row <= len(matrix)-1 && col < len(target) {
		if matrix[row][col] == target[col] {
			col++
			if col == len(target) {
				return row, true
			}
		} else {
			col = 0
			row++
			if row >= len(matrix) {
				return -1, false
			}
		}
	}
	return -1, false
}

func main(){

}