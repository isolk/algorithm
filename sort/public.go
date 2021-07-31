package sort

func swap(source []int, a, b int) {
	t := source[a]
	source[a] = source[b]
	source[b] = t
}
