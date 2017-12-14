package main

import (
	"algorithm/Model"
	"fmt"
)

func main() {
	fmt.Println("before Merge-sort:", Model.Arr)
	fmt.Println("after Merge-sort:", MergeSort(Model.Arr))
}

func MergeSort(arr[]int) []int {
	length := len(arr)
	if length <= 1 {
		 return arr
	}
	mid := length / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	return Merge(left, right)
}

func Merge(left, right []int) []int {
	l,r := 0,0
	result := []int{}

	for l < len(left) && r < len(right) {
		if l < len(left) && left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
