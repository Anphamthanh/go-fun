package main

import "testing"

func TestMedian1(t *testing.T) {
	median := NewMedian()
	in := [...]float64{3, 4, 5, 87, 23, 45, 123, 65, 23, 12, 78}
	ans := [...]float64{3, 3.5, 4, 4.5, 5, 14, 23, 34, 23, 23}
	for i, n := range ans {
		median.Add(in[i])
		if median.Get() != n {
			t.Errorf("Wrong value. %v Want %s got %s", median, n, median.Get())
		}
	}
}
