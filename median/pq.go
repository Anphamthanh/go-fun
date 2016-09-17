package main

import "errors"

type PQ struct {
	cmpFunc func(Item, Item) int
	heap    []Item
}

func NewPQ(cmp func(Item, Item) int) *PQ {
	return &PQ{
		cmpFunc: cmp,
		heap:    []Item{Item{}},
	}
}

func (pq PQ) isEmpty() bool {
	return pq.Size() == 0
}

func (pq PQ) Size() int {
	return len(pq.heap) - 1
}

func (pq *PQ) swapItem(i1, i2 int) {
	pq.heap[i1], pq.heap[i2] = pq.heap[i2], pq.heap[i1]
}

func (pq *PQ) maxItem(i1, i2 int) Item {
	if pq.cmpFunc(pq.heap[i1], pq.heap[i2]) > 0 {
		return pq.heap[i1]
	} else {
		return pq.heap[i2]
	}
}

func (pq *PQ) Push(i Item) error {
	pq.heap = append(pq.heap, i)
	n := len(pq.heap) - 1
	p := n / 2
	for p > 0 && pq.cmpFunc(pq.heap[p], pq.heap[n]) < 0 {
		pq.swapItem(p, n)
		n = p
		p = n / 2
	}
	return nil
}

func (pq *PQ) Pop() (Item, error) {
	if pq.Size() == 0 {
		return Item{}, errors.New("No more item")
	}
	pq.swapItem(0, len(pq.heap)-1)
	out := pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	//p = 1
	//c1 = 2
	//c2 = 3
	//for p < len(pq.heap) && pq.cmpFunc(pq.heap[p], pq.maxItem(c1, c2)) < 0 {
	//	}
	return out, nil
}
