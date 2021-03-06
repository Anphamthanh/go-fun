package main

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

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}

func (pq *PQ) swapItem(i1, i2 int) {
	if max(i1, i2) >= len(pq.heap) {
		panic("Swap index out of bound")
	}
	pq.heap[i1], pq.heap[i2] = pq.heap[i2], pq.heap[i1]
}

func (pq PQ) Peep() Item {
	if pq.Size() == 0 {
		panic("No thing to peep")
	}
	return pq.heap[1]
}

func (pq *PQ) Push(i Item) {
	pq.heap = append(pq.heap, i)
	n := len(pq.heap) - 1
	for p := n / 2; p > 0 && pq.cmpFunc(pq.heap[p], pq.heap[n]) < 0; {
		pq.swapItem(p, n)
		n, p = p, p/2
	}
}

func (pq *PQ) Pop() Item {
	if pq.Size() == 0 {
		panic("No thing to pop")
	}
	pq.swapItem(1, len(pq.heap)-1)
	out := pq.heap[len(pq.heap)-1]
	pq.heap = pq.heap[:len(pq.heap)-1]
	pq.fixDown(1)
	return out
}

func (pq *PQ) Remove(r Item) bool {
	for i, n := range pq.heap {
		if n.Equal(r) {
			pq.swapItem(i, len(pq.heap)-1)
			pq.heap = pq.heap[:len(pq.heap)-1]
			pq.fixDown(i)
			return true
		}
	}
	return false
}

func (pq *PQ) fixDown(p int) {
	for c := p * 2; p < len(pq.heap); {
		if c > len(pq.heap)-1 {
			break
		}
		if c+1 > len(pq.heap)-1 {
			if pq.cmpFunc(pq.heap[p], pq.heap[c]) < 0 {
				pq.swapItem(p, c)
			}
			break
		}
		if pq.cmpFunc(pq.heap[c+1], pq.heap[c]) > 0 {
			c += 1
		}
		if pq.cmpFunc(pq.heap[p], pq.heap[c]) < 0 {
			pq.swapItem(p, c)
			p, c = c, 2*c
		} else {
			break
		}
	}
}
