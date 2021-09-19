package components

func DirectInsertSort(src []int) {
	var temp int

	for i := 1; i < len(src); i++ {
		if src[i] < src[i-1] {
			temp = src[i]
			j := i - 1
			for ; j >= 0 && src[j] > temp; j-- {
				src[j+1] = src[j]
			}
			src[j+1] = temp
		}
	}
	return
}

func HarfInsertSort(src []int) {
	var low, mid, high int
	for i := 1; i < len(src); i++ {
		temp := src[i]
		low = 0
		high = i - 1
		for low <= high {
			mid = (low + high) / 2
			if src[mid] > temp {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		//now low = high + 1 and before low = high = mid
		for j := i - 1; j >= low; j-- {
			src[j+1] = src[j]
		}
		src[low] = temp
	}
}

func ShellSort(src []int, increment int) {
	//group
	for k := len(src) / increment; k >= 1; k = k / increment {
		//direcrinsetsort
		for i := k; i < len(src); i++ {
			if src[i] < src[i-k] {
				temp := src[i]
				j := i - k
				for ; j >= 0 && src[j] > temp; j = j - k {
					src[j+k] = src[j]
				}
				src[j+k] = temp
			}
		}
	}
}
