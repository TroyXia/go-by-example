package main

import (
	"container/heap"
	"fmt"
)

// Item 表示优先级队列中的元素
type Item struct {
	value    string // 元素的值
	priority int    // 元素的优先级
	index    int    // 元素在堆中的索引
}

// PriorityQueue 实现了 heap.Interface 接口
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 这里定义优先级高的元素先出队
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

// Push 向优先级队列中添加元素
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

// Pop 从优先级队列中移除并返回优先级最高的元素
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // 避免内存泄漏
	item.index = -1 // 表示该元素已不在队列中
	*pq = old[0 : n-1]
	return item
}

// update 修改队列中元素的优先级和值
func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func main() {
	// 创建一个空的优先级队列
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 向队列中添加元素
	items := []*Item{
		{"item1", 1, 0},
		{"item2", 3, 0},
		{"item3", 2, 0},
	}
	for _, item := range items {
		heap.Push(&pq, item)
	}

	// 修改元素的优先级
	item := pq[0]
	pq.update(item, item.value, 0)

	// 依次取出队列中的元素
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		fmt.Printf("Value: %s, Priority: %d\n", item.value, item.priority)
	}
}
