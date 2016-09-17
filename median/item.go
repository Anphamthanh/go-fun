package main

type Item struct {
	value    interface{}
	priority float64
}

func NewItem(value float64) *Item {
	return &Item{value: value, priority: value}
}
