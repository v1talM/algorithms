package main

import (
	"algorithm/Model"
	"fmt"
)

func main() {
	fmt.Println("before Selection-sort:", Model.Arr)
	fmt.Println("after Selection-sort:", SelectionSort(Model.Arr))
}

func SelectionSort(arr[]int) []int {
	for i := 0; i < Model.Len - 1; i++ {
		min := i
		for j := i+1; j <Model.Len; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			arr[min], arr[i] = arr[i], arr[min]
		}
	}
	return arr
}
