package sort

func Bucket(source []int, min, max int) {
	if min >= max {
		return
	}

	span := 0
	if min < 0 {
		span = -min
	}

	bucketsCount := (max-min)/10 + 1
	buckets := make([][]int, int(bucketsCount))

	for _, v := range source {
		index := (v + span) / int(bucketsCount)
		buckets[index] = append(buckets[index], v)
	}

	for _, v := range buckets {
		MergerSort(v)
	}

	i := 0
	for _, v := range buckets {
		for _, y := range v {
			source[i] = y
			i++
		}
	}
}
