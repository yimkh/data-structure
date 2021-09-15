# data-structure
## sqList 线性表
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

func (sqList SqLists) PrintSqList()

func (sqList SqLists) PrintLength()

func (sqList SqLists) IsEmpty() bool

func (sqList SqLists) IsFull() bool 

func (sqList *SqLists) Add(data int)

func (sqList *SqLists) AddList(data []int)

func (sqList *SqLists) Insert(locat int, data int) 

func (sqList *SqLists) Delete(locat int)

func (sqList SqLists) GetElem(locat int) int

func (sqList SqLists) LocateElem(value int) int
```
## 单链表
规定：  <br>
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