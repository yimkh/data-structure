package struction

import "fmt"

var orders struct {
	maxSize int
	form    []int
	length  int
}

func Init(maxSize int) {
	orders.maxSize = maxSize
	orders.form = make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		orders.form[i] = 0
	}
	orders.length = 0
}

func PrintOrder() {
	for i := 0; i < orders.maxSize; i++ {
		fmt.Printf("form[%v] = %v\n", i, orders.form[i])
	}
}
