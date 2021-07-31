package search

func Binary(source []int, val int) int {
	if len(source) == 0 {
		return -1
	}
	return binary(source, 0, len(source)-1, val)
}

func BinaryFirst(source []int, val int) int {
	if len(source) == 0 {
		return -1
	}
	return binaryFirst(source, 0, len(source)-1, val)
}

func BinaryLast(source []int, val int) int {
	if len(source) == 0 {
		return -1
	}
	return binaryLast2(source, 0, len(source)-1, val)
}

func BinaryFirstBig(source []int, val int) int {
	if len(source) == 0 {
		return -1
	}
	return binaryFirstBig(source, 0, len(source)-1, val)
}

func binary(source []int, left, right, val int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if source[mid] == val {
		return mid
	} else if val < source[mid] {
		return binary(source, left, mid-1, val)
	} else {
		return binary(source, mid+1, right, val)
	}
}

// 注意退出条件
func binaryFirst(source []int, left, right, val int) int {
	if right < left {
		return -1
	}

	mid := left + (right-left)/2
	if source[mid] >= val {
		if source[mid] == val && (mid == 0 || source[mid-1] != val) {
			return mid
		}
		return binaryFirst(source, left, mid-1, val)
	}

	return binaryFirst(source, mid+1, right, val)
}

func binaryLast2(source []int, left, right, val int) int {
	if left > right {
		return -1
	}
	if left == right {
		if source[left] == val {
			return left
		}
		return -1
	}

	mid := left + (right-left)/2 + 1
	if source[mid] == val {
		return binaryLast(source, mid, right, val)
	} else if val < source[mid] {
		return binaryLast(source, left, mid-1, val)
	} else {
		return binaryLast(source, mid+1, right, val)
	}
}

func binaryLast(source []int, left, right, val int) int {
	if left > right {
		return -1
	}
	if left == right {
		if source[left] == val {
			return left
		}
		return -1
	}

	mid := left + (right-left)/2 + 1
	if source[mid] == val {
		return binaryLast(source, mid, right, val)
	} else if val < source[mid] {
		return binaryLast(source, left, mid-1, val)
	} else {
		return binaryLast(source, mid+1, right, val)
	}
}

// 查找第一个大于等于的数
// 首先正常二分，如果这个数大于等于目标数那么：
//   如果如果当前是第一个数，那么就返回
//   如果不是第一个数，但是前一个数比当前数小，返回
//   否则，查左边。
// 否则，查右边。
// 重点是要掌握等于和大于的关系
func binaryFirstBig(source []int, left, right, val int) int {
	if left > right {
		return -1
	}

	mid := (right + left) / 2
	if source[mid] >= val {
		if mid == 0 || source[mid-1] < val {
			return mid
		} else {
			return binaryFirstBig(source, left, mid-1, val)
		}
	}

	return binaryFirstBig(source, mid+1, right, val)
}
