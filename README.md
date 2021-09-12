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
### 单链表
规定：
1.带头结点
2.每个真正结点的值均 >= 0
3.增删改查按顺序，1代表第一个有真正值的结点（头结点次序为0，值为-1）
4.返回-1，代表没有此项值或结点次序（头结点除外）