package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var (
	cmp1 = func(i1 Item, i2 Item) int {
		i := i1.priority - i2.priority
		if i == 0 {
			return 0
		} else {
			return int(i / math.Abs(i))
		}
	}
	cmp2 = func(i1 Item, i2 Item) int {
		return -cmp1(i1, i2)
	}
	median = NewMedian()
)

func NewMedian() *Median {
	return &Median{
		lower: NewPQ(cmp1),
		upper: NewPQ(cmp2),
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	input := []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	N, err := strconv.Atoi(input[0])
	if err != nil {
		panic("Error parsing " + input[0])
	}
	input = input[1:]
	fmt.Println(N)
	fmt.Println(input)
	median.Add(7)
	fmt.Println(median)
}
