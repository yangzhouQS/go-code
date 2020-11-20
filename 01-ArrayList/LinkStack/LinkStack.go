package LinkStack

// 链式栈
type Node struct {
	data  interface{}
	pNext *Node // 存储下一个节点
}

type LinkStack interface {
	IsEmpty() bool
	Push(value interface{})
	Pop() interface{}
	Length() int
}
