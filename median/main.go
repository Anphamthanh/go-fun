package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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
	input = input[1:]
	c, v := "", 0.0
	for _, cmd := range input {
		c = strings.Split(cmd, " ")[0]
		v, _ = strconv.ParseFloat(strings.Split(cmd, " ")[1], 64)
		if c == "r" {
			if median.Remove(v) {
				if median.lower.Size()+median.upper.Size() == 0 {
					fmt.Println("Wrong!")
				} else {
					fmt.Println(median.Get())
				}
			} else {
				fmt.Println("Wrong!")
			}
		} else {
			median.Add(v)
			fmt.Println(median.Get())
		}
	}
}
