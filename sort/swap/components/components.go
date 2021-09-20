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

func QuickSort(src []int) {
	quickSort(src, 0, len(src)-1)
}

func quickSort(src []int, low int, high int) {
	if low >= high {
		return
	}
	pivotpos := Partition(src, low, high)
	quickSort(src, low, pivotpos-1)
	quickSort(src, pivotpos+1, high)
}

func Partition(src []int, low int, high int) (pivotpos int) {
	pivotposv := src[low]
	for low < high {
		for low < high && src[high] >= pivotposv {
			high--
		}
		src[low] = src[high]
		for low < high && src[low] <= pivotposv {
			low++
		}
		src[high] = src[low]
	}
	src[low] = pivotposv
	pivotpos = low
	return pivotpos
}
