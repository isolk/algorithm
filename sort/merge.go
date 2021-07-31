package sort

func MergerSort(source []int) {
	if len(source) <= 1 {
		return
	}
	mid := len(source) / 2
	MergerSort(source[:mid])
	MergerSort(source[mid:])
	merge(source[:mid], source[mid:], source)
}

func merge(a []int, b []int, origin []int) {
	result := make([]int, len(a)+len(b))
	resultIndex := 0
	ia, ib := 0, 0
	for ia < len(a) && ib < len(b) {
		if a[ia] <= b[ib] { // 注意必须是小于等于，否则结果是不稳定的。merge时，相同元素放在前面可以保证不改变相对位置。
			result[resultIndex] = a[ia]
			ia++
		} else {
			result[resultIndex] = b[ib]
			ib++
		}
		resultIndex++
	}

	if ia < len(a) {
		copy(result[resultIndex:], a[ia:])
	} else if ib < len(b) {
		copy(result[resultIndex:], b[ib:])
	}
	copy(origin, result)
}
