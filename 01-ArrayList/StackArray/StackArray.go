package StackArray

type StackArray interface {
	Size() int        // 大小
	Clear()           // 清空
	Pop() interface{} // 弹出
	Push()            // 压栈
	IsFull() bool     // 是否满了
	IsEmpty() bool    // 是否为空
}
type Stack struct {
	dataSource  []interface{} // 保存数据容器
	capSize     int           // 最大容量
	currentSize int           // 当前大小

}

func NewStack() *Stack {
	stack := new(Stack)
	stack.dataSource = make([]interface{}, 0, 10)
	stack.capSize = 10
	stack.currentSize = 0
	return stack
}

// 栈大小
func (stack *Stack) Size() int {
	return stack.currentSize
}

// 清空栈
func (stack *Stack) Clear() {
	stack.dataSource = make([]interface{}, 0, 10)
	stack.currentSize = 10
	stack.currentSize = 0
}

// 出栈
func (stack *Stack) Pop() interface{} {
	if !stack.IsEmpty() {
		lastVal := stack.dataSource[stack.currentSize-1]
		stack.dataSource = stack.dataSource[0 : stack.currentSize-1]
		stack.currentSize--
		return lastVal // 返回最后一个数据
	}
	return nil
}

// 压栈
func (stack *Stack) Push(value interface{}) {
	// 栈满了
	if stack.IsFull() {

	} else { // 栈没有满
		stack.dataSource = append(stack.dataSource, value)
		stack.currentSize++
	}
}

// 栈是否满了
func (stack *Stack) IsFull() bool {
	return stack.currentSize >= stack.capSize
}

// 是否为空
func (stack *Stack) IsEmpty() bool {
	return stack.currentSize == 0
}
