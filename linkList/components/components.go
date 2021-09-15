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

func IsEmpty(lNode LNode) bool {
	if lNode.NextlNode == nil {
		fmt.Println("empty!!!")
		return true
	} else {
		return false
	}
}

func IsIllegal(lNode *LNode) bool {
	if lNode == nil {
		fmt.Println("lnode is nil")
		return true
	}
	return false
}

func Length(lNode *LNode) int {
	var lNodeScanner *LNode
	lNodeScanner = lNode
	i := 0
	for ; lNodeScanner.NextlNode != nil; i++ {
		lNodeScanner = lNodeScanner.NextlNode
	}
	if i == 0 {
		return 0
	}
	return i
}

func PrintlNode(lNode LNode) {
	if IsIllegal(&lNode) {
		return
	}
	var lNodeScanner *LNode
	lNodeScanner = &lNode
	if IsEmpty(*lNodeScanner) {
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

func AddInHead(lNode *LNode, newlNode *LNode) {
	if IsIllegal(newlNode) {
		return
	}
	var lNodeScanner *LNode
	lNodeScanner = lNode
	newlNode.NextlNode = lNodeScanner.NextlNode
	lNodeScanner.NextlNode = newlNode
}

func AddInEnd(lNode *LNode, newlNode *LNode) {
	if IsIllegal(newlNode) {
		return
	}
	var lNodeScanner *LNode
	lNodeScanner = lNode
	for lNodeScanner.NextlNode != nil {
		lNodeScanner = lNodeScanner.NextlNode
	}
	newlNode.NextlNode = lNodeScanner.NextlNode
	lNodeScanner.NextlNode = newlNode
}

func Delete(lNode *LNode, locat int) (elem int) {
	priorNode := GetElem(lNode, locat-1)
	if IsIllegal(priorNode) {
		return -1
	}
	elem = priorNode.NextlNode.Data
	priorNode.NextlNode = priorNode.NextlNode.NextlNode
	return elem
}

func DeleteNode(lNode *LNode, deleteNode *LNode) {
	deleteNode.Data = deleteNode.NextlNode.Data
	deleteNode.NextlNode = deleteNode.NextlNode.NextlNode
}

func Insert(lNode *LNode, locat int, newlNode *LNode) {
	if IsIllegal(newlNode) {
		return
	}
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

func InsertNext(lNode *LNode, newlNode *LNode) {
	if IsIllegal(newlNode) {
		return
	}
	newlNode.NextlNode = lNode.NextlNode
	lNode.NextlNode = newlNode
}

func InsertPrior(lNode *LNode, newlNode *LNode) {
	if IsIllegal(newlNode) {
		return
	}
	newlNode.NextlNode = lNode.NextlNode
	lNode.NextlNode = newlNode
	temp := newlNode.Data
	newlNode.Data = lNode.Data
	lNode.Data = temp
	return
}

func GetElem(lNode *LNode, locat int) *LNode {
	if locat < 1 {
		fmt.Println("locat is illegal")
		return Init().NextlNode
	} else {
		var lNodeScanner *LNode
		lNodeScanner = lNode
		i := 0
		for ; i < locat && lNodeScanner.NextlNode != nil; i++ {
			lNodeScanner = lNodeScanner.NextlNode
		}
		if i < locat {
			fmt.Println("up to max")
			return Init().NextlNode
		}
		return lNodeScanner
	}
}

func LocatElem(lNode *LNode, value int) (*LNode, int) {
	var lNodeScanner *LNode
	lNodeScanner = lNode
	i := 0
	for ; lNodeScanner.NextlNode != nil; i++ {
		if lNodeScanner.Data == value {
			return lNodeScanner, i
		}
		lNodeScanner = lNodeScanner.NextlNode
	}
	return Init().NextlNode, -1
}
