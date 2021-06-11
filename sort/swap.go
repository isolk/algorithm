package sort

func Swap(ary []int, i, j int) {
	t := ary[i]
	ary[i] = ary[j]
	ary[j] = t
}
