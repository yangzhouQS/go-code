package Queue

type MyQueue interface {
	Size() int
	Front() interface{}                    // 第一个元素
	End() interface{}                      // 队列的尾部的元素
	Enqueue(value interface{}) interface{} // 入队操作
	Dequeue()                              // 出队操作
	IsEmpty() bool                         // 是否为空
	Clear()
}
type Queue struct {
	dataStrore []interface{}
	theSize    int
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.dataStrore = make([]interface{}, 0, 10)
	queue.theSize = 0
	return queue
}

func (queue *Queue) Size() int {
	return queue.theSize
}

func (queue *Queue) Front() interface{} {
	if queue.theSize == 0 {
		return nil
	} else {
		return queue.dataStrore[0]
	}
}

func (queue *Queue) End() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	return queue.dataStrore[queue.Size()-1]
}

// 入队操作
func (queue *Queue) Enqueue(value interface{}) {
	queue.dataStrore = append(queue.dataStrore, value)
	queue.theSize++
}

// 出队操作
func (queue *Queue) Dequeue() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	data := queue.dataStrore[0]
	if queue.Size() > 1 {
		queue.dataStrore = queue.dataStrore[1:queue.Size()]
	} else {
		queue.dataStrore = make([]interface{}, 0)
	}
	queue.theSize--
	return data
}

// 判断是否为空
func (queue *Queue) IsEmpty() bool {
	return queue.theSize == 0
}
func (queue *Queue) Clear() {
	queue.dataStrore = make([]interface{}, 0, 10)
	queue.theSize = 0
}
