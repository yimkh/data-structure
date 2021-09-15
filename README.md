# data-structure
## 目录
1.线性表  <br>
2.单链表  <br>
## sqList 线性表
### 规定: 
1.顺序表从0开始，次序从1开始  <br>
2.元素的值均 >= 0, 次序均 >= 1  <br>
### 结构体
```
type SqLists struct {
	maxSize int
	data    []int
	length  int
}
```
### 方法
```
func Init(maxSize int) SqLists

func PrintSqList(sqList SqLists)

func PrintLength(sqList SqLists)

func IsEmpty(sqList SqLists) bool

func IsFull(sqList SqLists) bool 

func Add(sqList *SqLists, data int) 

func AddList(sqList *SqLists, data []int) 

func Delete(sqList *SqLists, locat int) 

func Insert(sqList *SqLists, locat int, data int) 

func GetElem(sqList SqLists, locat int) int 

func LocateElem(sqList SqLists, value int) int 


```
## 单链表
### 规定:
1.带头结点（头结点次序为0，值为-1） <br>
2.每个结点除头结点外，次序和值均 >= 0  <br>
3.增删改查按次序，1代表第一个有真正值的结点  <br>
4.返回-1，代表没有此项值或结点次序（头结点除外）  <br>
### 结构体
```
type LNode struct {
	Data      int
	NextlNode *LNode
}
```
### 方法
```
func Init() LNode

func IsEmpty(lNode LNode) bool

func IsIllegal(lNode *LNode) bool

func Length(lNode *LNode) int

func PrintlNode(lNode LNode)

func AddInHead(lNode *LNode, newlNode *LNode)

func AddInEnd(lNode *LNode, newlNode *LNode)

func Delete(lNode *LNode, locat int) (elem int)

func DeleteNode(lNode *LNode, deleteNode *LNode)

func Insert(lNode *LNode, locat int, newlNode *LNode)

func InsertNext(lNode *LNode, newlNode *LNode)

func InsertPrior(lNode *LNode, newlNode *LNode)

func GetElem(lNode *LNode, locat int) *LNode

func LocatElem(lNode *LNode, value int)
```