package Queue

type Queue struct {
	dataStore []interface{}
	theSize   int
}

func NewQueue() *Queue {
	return &Queue{dataStore: make([]interface{}, 0), theSize: 0}
}

type MyQueue interface {
	Front() interface{}
	End() interface{}
	Size() int
	Clear()
	IsEmpty() bool
	EnQueue(value interface{})
	DeQueue() interface{}
}

func (queue *Queue) Front() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	return queue.dataStore[0]
}

func (queue *Queue) End() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	return queue.dataStore[queue.Size()-1]
}

func (queue *Queue) Size() int {
	return queue.theSize
}

func (queue *Queue) Clear() {
	queue.dataStore = make([]interface{}, 0)
	queue.theSize = 0
}

func (queue *Queue) IsEmpty() bool {
	return queue.Size() == 0
}

func (queue *Queue) EnQueue(value interface{}) {
	queue.dataStore = append(queue.dataStore, value)
	queue.theSize++
}

func (queue *Queue) DeQueue() interface{} {
	if queue.Size() == 0 {
		return nil
	}
	data := queue.dataStore[0]
	if queue.Size() > 1 {
		queue.dataStore = queue.dataStore[1:queue.Size()]
	} else {
		queue.dataStore = make([]interface{}, 0)
	}
	queue.theSize--
	return data
}
