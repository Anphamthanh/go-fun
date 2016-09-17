package main

import (
	"fmt"
	"testing"
)

func TestEmpty1(t *testing.T) {
	pq := NewPQ(cmp)
	if !pq.isEmpty() {
		t.Errorf("Should be empty!")
	}

	pq.Push(Item{value: 2, priority: 2})
	if pq.isEmpty() {
		t.Errorf("Should not be empty!")
	}
}

func TestPush1(t *testing.T) {
	pq := NewPQ(cmp)
	items := [...]float64{1, 2, 3, 4, 5, 6}
	for _, item := range items {
		pq.Push(*NewItem(item))
	}
	exp := "[{<nil> 0} {6 6} {4 4} {5 5} {1 1} {3 3} {2 2}]"
	if fmt.Sprintf("%v", pq.heap) != exp {
		t.Errorf("%v", pq.heap)
	}
	if pq.Size() != 6 {
		t.Errorf("Expect size to be 6")
	}
}

func TestPush2(t *testing.T) {
	pq := NewPQ(cmp)
	items := [...]float64{4, 1, 2, 3, 5, 6}
	for _, item := range items {
		pq.Push(Item{value: item, priority: -item})
	}
	exp := "[{<nil> 0} {1 -1} {3 -3} {2 -2} {4 -4} {5 -5} {6 -6}]"
	if fmt.Sprintf("%v", pq.heap) != exp {
		t.Errorf("%v", pq.heap)
	}
	if pq.Size() != 6 {
		t.Errorf("Expect size to be 6")
	}
}

func TestPop1(t *testing.T) {
	pq := NewPQ(cmp)
	_, err := pq.Pop()
	if err == nil {
		t.Errorf("Nothing to pop")
	}
	items := [...]float64{4, 1, 2, 3, 6, 5}
	for _, item := range items {
		pq.Push(*NewItem(item))
	}
	if fmt.Sprintf("%v", pq.heap) != "[{<nil> 0} {6 6} {4 4} {5 5} {1 1} {3 3} {2 2}]" {
		t.Errorf("%v", pq.heap)
	}
	answers := [...]float64{6, 5, 4, 3, 2, 1, 0}
	for _, ans := range answers {
		item, err := pq.Pop()
		if err != nil {
			if pq.Size() == 0 {
				continue
			} else {
				t.Errorf("Unexpected %v", pq.heap)
			}
		}
		if item.value != ans {
			t.Errorf("Expect %s, got %s, pq %v", ans, item, pq.heap)
		}
	}
}
