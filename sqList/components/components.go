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

func PrintSqList(sqList SqLists) {
	if IsEmpty(sqList) {
		fmt.Println("empty!!!")
	} else {
		for i := 0; i < sqList.length; i++ {
			fmt.Printf("data[%v] = %v\n", i, sqList.data[i])
		}
	}
}

func PrintLength(sqList SqLists) {
	fmt.Printf("length = %v\n", sqList.length)
}

func IsEmpty(sqList SqLists) bool {
	if sqList.length == 0 {
		return true
	} else {
		return false
	}
}

func IsFull(sqList SqLists) bool {
	if sqList.length >= sqList.maxSize {
		return true
	} else {
		return false
	}
}

func Add(sqList *SqLists, data int) {
	if IsFull(*sqList) {
		fmt.Println("full!!!")
	} else {
		sqList.data[sqList.length] = data
		sqList.length++
	}
}

func AddList(sqList *SqLists, data []int) {
	if IsFull(*sqList) {
		fmt.Println("full!!!")
	} else if sqList.length+len(data) > sqList.maxSize {
		fmt.Println("up to max")
	} else {
		for i := 0; i < len(data); i++ {
			Add(sqList, data[i])
		}
	}
}

func Delete(sqList *SqLists, locat int) {
	if locat < 1 || locat > sqList.length {
		fmt.Println("locat is illegal")
	} else if IsEmpty(*sqList) {
		fmt.Println("empty!!!")
	} else {
		for i := locat - 1; i < sqList.length-1; i++ {
			sqList.data[i] = sqList.data[i+1]
		}
		sqList.data[sqList.length-1] = 0
		sqList.length--
	}
}

func Insert(sqList *SqLists, locat int, data int) {
	if locat < 1 || locat > sqList.length+1 {
		fmt.Println("locat is illegal")
	} else if IsFull(*sqList) {
		fmt.Println("full!!!")
	} else {
		for i := sqList.length; i >= locat; i-- {
			sqList.data[i] = sqList.data[i-1]
		}
		sqList.data[locat] = data
		sqList.length++
	}
}

func GetElem(sqList SqLists, locat int) int {
	if locat < 1 || locat > sqList.length {
		fmt.Println("locat is illegal")
		return 0
	} else if IsEmpty(sqList) {
		fmt.Println("empty!!!")
		return 0
	} else {
		return sqList.data[locat-1]
	}
}

func LocateElem(sqList SqLists, value int) int {
	for i := 0; i < sqList.length; i++ {
		if sqList.data[i] == value {
			return i + 1
		}
	}
	return 0
}
