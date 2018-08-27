package lintcode

import (
	"container/heap"
	"fmt"
	"sort"
)

func topk(nums []int, k int) []int { //lintcode 544
	// write your code here
	pq := make(PriorityQueueInt, k)
	for key, value := range nums {
		if key < k {
			pq[key] = &intItem{
				value:    value,
				priority: value * -1,
				index:    key,
			}
		} else {
			if key == k {
				heap.Init(&pq)
			}
			if pq.look().(*intItem).priority < value {
				pq[0] = &intItem{value, value * -1, 0}
				heap.Fix(&pq, 0)
			}
		}
	}
	if len(nums) == k {
		heap.Init(&pq)
	}
	ret := make([]int, k)
	for i := k - 1; i >= 0; i-- {
		ret[i] = heap.Pop(&pq).(*intItem).value
	}
	return ret[:]
}

func topKFrequentWords(words []string, k int) []string { //lintcode 471
	// write your code here
	pq := make(PriorityQueue, k)
	items := map[string]int{}
	for _, value := range words {
		_, ok := items[value]
		if ok {
			items[value] = items[value] + 1
		} else {
			items[value] = 1
		}
	}
	fmt.Println(items)
	i := 0
	for word, freq := range items {
		if i < k {
			pq[i] = &stringItem{
				value:    word,
				priority: freq * -1,
				index:    i,
			}
		} else {
			if i == k {
				heap.Init(&pq)
			}
			fmt.Println(pq[0].value, pq[0].priority, freq*-1)
			//if pq.look().(*stringItem).priority == freq*-1 {
			if pq[0].priority == freq*-1 {
				fmt.Println("============")
				str := []string{pq[0].value, word}
				sort.Strings(str)
				fmt.Println(str)
				pq[0] = &stringItem{str[0], freq * -1, 0}
			}
			//if pq.look().(*stringItem).priority > freq* -1 {
			if pq[0].priority > freq*-1 {
				fmt.Println("ccccccccccccc")
				pq[0] = &stringItem{word, freq * -1, 0}
				heap.Fix(&pq, 0)
			}
		}
		i++
	}

	if len(items) == k {
		heap.Init(&pq)
	}
	ret := make([]string, k)
	for i := k - 1; i >= 0; i-- {
		ret[i] = heap.Pop(&pq).(*stringItem).value
	}
	return ret[:]
}

// This example demonstrates a priority queue built using the heap interface.

// An stringItem is something we manage in a priority queue.
type stringItem struct {
	value    string // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*stringItem

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*stringItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq PriorityQueue) look() interface{} {
	item := pq[0]
	return item
}

// update modifies the priority and value of an stringItem in the queue.
func (pq *PriorityQueue) update(item *stringItem, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

type intItem struct {
	value    int // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueueInt []*intItem

func (pq PriorityQueueInt) Len() int { return len(pq) }
func (pq PriorityQueueInt) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority > pq[j].priority
}
func (pq PriorityQueueInt) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueueInt) Push(x interface{}) {
	n := len(*pq)
	item := x.(*intItem)
	item.index = n
	*pq = append(*pq, item)
}
func (pq *PriorityQueueInt) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
func (pq PriorityQueueInt) look() interface{} {
	return pq[0]
}

// update modifies the priority and value of an stringItem in the queue.
func (pq *PriorityQueueInt) update(item *intItem, value int, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
func ExamplePriorityQueue() {
	// Some items and their priorities.
	items := map[string]int{
		"banana": 3, "apple": 2, "pear": 4,
	}
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	pq := make(PriorityQueue, len(items))
	i := 0
	for value, priority := range items {
		pq[i] = &stringItem{
			value:    value,
			priority: priority,
			index:    i,
		}
		i++
	}
	heap.Init(&pq)
	// Insert a new item and then modify its priority.
	item := &stringItem{
		value:    "orange",
		priority: 1,
	}
	heap.Push(&pq, item)
	pq.update(item, item.value, 5)
	// Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*stringItem)
		fmt.Printf("%.2d:%s ", item.priority, item.value)
	}
	// Output:
	// 05:orange 04:pear 03:banana 02:apple
}
