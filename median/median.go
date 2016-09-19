package main

import "fmt"

type Median struct {
	lower *PQ
	upper *PQ
}

func (m Median) String() string {
	return fmt.Sprintf("lower: %v | upper: %v", m.lower, m.upper)
}

func (m *Median) Get() float64 {
	if m.lower.Size() == 0 {
		panic("Empty list, no median")
	}
	if m.lower.Size() == m.upper.Size()+1 {
		return m.lower.Peep().value.(float64)
	}
	if m.lower.Size() == m.upper.Size() {
		return (m.lower.Peep().value.(float64) + m.upper.Peep().value.(float64)) / 2
	}
	panic("Unexpected situation")
}

func (m *Median) Add(f float64) {
	item := Item{value: f, priority: f}
	if m.lower.Size() == m.upper.Size() {
		if m.upper.Size() == 0 {
			m.lower.Push(item)
		} else {
			if f < m.upper.Peep().value.(float64) {
				m.lower.Push(item)
			} else {
				m.lower.Push(m.upper.Pop())
				m.upper.Push(item)
			}
		}
		return
	}
	if m.lower.Size() == m.upper.Size()+1 {
		if f > m.lower.Peep().value.(float64) {
			m.upper.Push(item)
		} else {
			m.upper.Push(m.lower.Pop())
			m.lower.Push(item)
		}
		return
	}
	panic("Unexpected case " + m.String())
}

func (m *Median) Remove(f float64) bool {
	item := Item{value: f, priority: f}
	if !m.lower.Remove(item) {
		return m.upper.Remove(item)
	}
	return true
}
