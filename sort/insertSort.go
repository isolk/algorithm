package sort

func InsertSort(ary []int) {
	for i := 0; i < len(ary)-1; i++ {
		value := ary[i+1]
		j := i
		for ; j >= 0; j-- {
			if value < ary[j] {
				ary[j+1] = ary[j]
			} else {
				break
			}
		}
		ary[j+1] = value
	}
}
