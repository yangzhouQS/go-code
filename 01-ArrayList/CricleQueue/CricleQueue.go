package CricleQueue

import "errors"

const QueueSize = 4 // 最多存储 QueueSize - 1个元素

type CricleQueue struct {
	dataStore [QueueSize]interface{} // 存储数据的结构
	front     int                    // 头部的位置
	rear      int                    // 尾部的位置
}

// 初始化
func InitQueue(c *CricleQueue) {
	c.front = 0
	c.rear = 0
}

func QueueLength(c *CricleQueue) int {
	return (c.rear - c.front + QueueSize) % QueueSize
}

// 入队
func EnQueue(c *CricleQueue, data interface{}) (error error) {
	if (c.rear+1)%QueueSize == c.front%QueueSize {
		return errors.New("队列满了")
	}
	c.dataStore[c.rear] = data        // 入队
	c.rear = (c.rear + 1) % QueueSize // 循环 , 尾部向后移动1
	return nil
}

// 出队
func DeQueue(c *CricleQueue) (data interface{}, err error) {
	if c.front == c.rear {
		return nil, errors.New("队列为空")
	}
	result := c.dataStore[c.front]      // 取出第一个数据
	c.front = (c.front + 1) % QueueSize // 101===>回到1, 头部向后移动
	return result, nil
}
