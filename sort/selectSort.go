package sort

func SelectSort(ary []int) {
	for i := 0; i < len(ary); i++ {
		minIndex := i
		for j := i + 1; j < len(ary); j++ {
			if ary[j] < ary[minIndex] {
				minIndex = j
			}
		}
		Swap(ary, minIndex, i)
	}
}
