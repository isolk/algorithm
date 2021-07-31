package sort
func Counting(source []int,min,max int){
	ary := make([]int,max-min)
	for _, value := range source {
		ary[value]+=1
	}

	// ary= 10,1,5,2
	count := 0
	for index, val := range ary {
		for i := 0; i < val; i++ {
			source[i+count]= index
		}
		count += val
	}
}
