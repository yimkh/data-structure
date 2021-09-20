package components

func BubbleSort(src []int) {
	flag := false

	for i := 0; i < len(src); i++ {
		for j := len(src) - 1; j > i; j-- {
			if src[j] < src[j-1] {
				temp := src[j]
				src[j] = src[j-1]
				src[j-1] = temp
				flag = true
			}
		}
		if !flag {
			return
		} else {
			flag = false
		}
	}

	return
}
