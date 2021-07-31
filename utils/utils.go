package utils

func SwapAry(ary []int, indexA, indxeB int) {
	t := ary[indexA]
	ary[indexA] = ary[indxeB]
	ary[indxeB] = t
}

func Max(ary []int, indexA, indexB int) int {
	if ary[indexA] > indexB {
		return indexA
	}
	return indexB
}
