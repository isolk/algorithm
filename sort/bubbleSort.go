package sort

func BubbleSort(ary []int) {
	for i := 0; i < len(ary); i++ {
		swaped := false
		for j := 0; j < len(ary)-i-1; j++ {
			if ary[j+1] < ary[j] {
				t := ary[j]
				ary[j] = ary[j+1]
				ary[j+1] = t
				swaped = true
			}
		}
		if !swaped {
			return
		}
	}
}
