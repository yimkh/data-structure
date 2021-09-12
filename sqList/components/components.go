package components

import "fmt"

type SqLists struct {
	maxSize int
	data    []int
	length  int
}

func Init(maxSize int) SqLists {
	var sqList SqLists

	sqList.maxSize = maxSize
	sqList.data = make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		sqList.data[i] = 0
	}
	sqList.length = 0

	return sqList
}

func (sqList SqLists) PrintSqList() {
	if sqList.IsEmpty() {
		fmt.Println("empty!!!")
	} else {
		for i := 0; i < sqList.length; i++ {
			fmt.Printf("data[%v] = %v\n", i, sqList.data[i])
		}
	}
}

func (sqList SqLists) PrintLength() {
	fmt.Printf("length = %v\n", sqList.length)
}

func (sqList SqLists) IsEmpty() bool {
	if sqList.length == 0 {
		return true
	} else {
		return false
	}
}

func (sqList SqLists) IsFull() bool {
	if sqList.length >= sqList.maxSize {
		return true
	} else {
		return false
	}
}

func (sqList *SqLists) Add(data int) {
	if sqList.IsFull() {
		fmt.Println("full!!!")
	} else {
		sqList.data[sqList.length] = data
		sqList.length++
	}
}

func (sqList *SqLists) AddList(data []int) {
	if sqList.IsFull() {
		fmt.Println("full!!!")
	} else if sqList.length+len(data) > sqList.maxSize {
		fmt.Println("up to max")
	} else {
		for i := 0; i < len(data); i++ {
			sqList.Add(data[i])
		}
	}
}

func (sqList *SqLists) Insert(locat int, data int) {
	if locat < 1 || locat > sqList.length+1 {
		fmt.Println("locat is illegal")
	} else if sqList.IsFull() {
		fmt.Println("full!!!")
	} else {
		for i := sqList.length; i >= locat; i-- {
			sqList.data[i] = sqList.data[i-1]
		}
		sqList.data[locat] = data
		sqList.length++
		sqList.PrintSqList()
	}
}

func (sqList SqLists) Delete(locat int) {
	if locat < 1 || locat > sqList.length {
		fmt.Println("locat is illegal")
	} else if sqList.IsEmpty() {
		fmt.Println("empty!!!")
	} else {
		for i := locat - 1; i < sqList.length-1; i++ {
			sqList.data[i] = sqList.data[i+1]
		}
		sqList.data[sqList.length-1] = 0
		sqList.length--
		sqList.PrintSqList()
	}
}
