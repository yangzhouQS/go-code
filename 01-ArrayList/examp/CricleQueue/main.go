package main

import (
	"01-ArrayList/CricleQueue"
	"fmt"
)

func main() {
	var queue CricleQueue.CricleQueue
	CricleQueue.InitQueue(&queue)
	CricleQueue.EnQueue(&queue, 1)
	CricleQueue.EnQueue(&queue, 2)
	CricleQueue.EnQueue(&queue, 3)
	CricleQueue.EnQueue(&queue, 4)
	CricleQueue.EnQueue(&queue, 5)
	fmt.Println(CricleQueue.DeQueue(&queue))
	fmt.Println(CricleQueue.DeQueue(&queue))
	fmt.Println(CricleQueue.DeQueue(&queue))
	fmt.Println(CricleQueue.DeQueue(&queue))
}
