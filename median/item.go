package main

type Item struct {
	value    interface{}
	priority float64
}

func NewItem(value float64) *Item {
	return &Item{value: value, priority: value}
}

func (i1 Item) Equal(i2 Item) bool {
	return i1.value == i2.value
}
