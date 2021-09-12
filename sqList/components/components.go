package components

import "fmt"

var sqList struct {
	maxSize int
	data    []int
	length  int
}

func Init(maxSize int) {
	sqList.maxSize = maxSize
	sqList.data = make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		sqList.data[i] = 0
	}
	sqList.length = 0
}

func PrintSqList() {
	if IsEmpty() {
		fmt.Println("empty!!!")
	} else {
		for i := 0; i < sqList.length; i++ {
			fmt.Printf("data[%v] = %v\n", i, sqList.data[i])
		}
	}
}

func IsEmpty() bool {
	if sqList.length == 0 {
		return true
	} else {
		return false
	}
}

func IsFull() bool {
	if sqList.length >= sqList.maxSize {
		return true
	} else {
		return false
	}
}

func Add(data int) {
	if IsFull() {
		fmt.Println("full!!!")
	} else {
		sqList.data[sqList.length] = data
		sqList.length++
		PrintSqList()
	}
}

func Insert(locat int, data int) {
	if locat < 1 || locat > sqList.length+1 {
		fmt.Println("locat is illegal")
	} else if IsFull() {
		fmt.Println("full!!!")
	} else {
		for i := sqList.length; i >= locat; i-- {
			sqList.data[i] = sqList.data[i-1]
		}
		sqList.data[locat] = data
		sqList.length++
		PrintSqList()
	}
}
