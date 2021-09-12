package components

import "fmt"

type LNode struct {
	Data      int
	NextlNode *LNode
}

func Init() LNode {
	return LNode{
		Data:      -1,
		NextlNode: nil,
	}
}

func (lNode LNode) IsEmpty() bool {
	if lNode.NextlNode == nil {
		fmt.Println("empty!!!")
		return true
	} else {
		return false
	}
}

func (lNode *LNode) AddInHead(newlNode *LNode) {
	if newlNode == nil {
		fmt.Println("lnode is nil")
		return
	}
	var lNodeScanner *LNode
	lNodeScanner = lNode
	newlNode.NextlNode = lNodeScanner.NextlNode
	lNodeScanner.NextlNode = newlNode
}

func (lNode LNode) PrintlNode() {
	var lNodeScanner *LNode
	lNodeScanner = &lNode
	if lNodeScanner.IsEmpty() {
		return
	}
	for i := 0; lNodeScanner != nil; i++ {
		if lNodeScanner.Data == -1 {
			lNodeScanner = lNodeScanner.NextlNode
		} else {
			fmt.Printf("%v th : %v\n", i, lNodeScanner.Data)
			lNodeScanner = lNodeScanner.NextlNode
		}
	}
}

func (lNode *LNode) Insert(locat int, newlNode *LNode) {
	if locat < 1 {
		fmt.Println("locat is illegal")
		return
	} else {
		var lNodeScanner *LNode
		lNodeScanner = lNode
		i := 0
		for ; i < locat && lNodeScanner.NextlNode != nil; i++ {
			lNodeScanner = lNodeScanner.NextlNode
		}
		if i < locat {
			fmt.Println("up to max")
			return
		}
		newlNode.NextlNode = lNodeScanner.NextlNode
		lNodeScanner.NextlNode = newlNode
		return
	}
}

func (lNode LNode) GetElem(locat int) (value int) {
	if locat < 1 {
		fmt.Println("locat is illegal")
		return -1
	} else {
		var lNodeScanner *LNode
		lNodeScanner = &lNode
		i := 0
		for ; i <= locat && lNodeScanner.NextlNode != nil; i++ {
			lNodeScanner = lNodeScanner.NextlNode
		}
		if i < locat {
			fmt.Println("up to max")
			return -1
		}
		return lNodeScanner.Data
	}
}
