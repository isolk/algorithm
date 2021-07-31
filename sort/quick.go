package sort

import (
	"math/rand"
)

func QuickSort(source []int) {
	if len(source) <= 1 {
		return
	}

	pivot := partition(source)
	QuickSort(source[:pivot])
	QuickSort(source[pivot+1:])
}

func partition(source []int) int {
	right := len(source) - 1
	randPivot := rand.Int() % len(source)
	swap(source, randPivot, right)
	pivotValue := source[right]
	left := 0
	for left < right {
		for left < right {
			if source[left] > pivotValue {
				source[right] = source[left]
				break
			} else {
				left++
			}
		}
		for left < right {
			if source[right] <= pivotValue {
				source[left] = source[right]
				break
			} else {
				right--
			}
		}
	}
	source[left] = pivotValue
	return left
}
