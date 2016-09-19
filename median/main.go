package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	cmp1 = func(i1 Item, i2 Item) int {
		i := i1.priority - i2.priority
		return int(i)
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
	started := false
	c, v, cmd := "", 0.0, []string{}
	dict := make(map[float64]int)
	for scanner.Scan() {
		if !started {
			started = true
			continue
		}
		cmd = strings.Split(scanner.Text(), " ")
		c = cmd[0]
		v, _ = strconv.ParseFloat(cmd[1], 64)
		if c == "r" {
			if _, ok := dict[v]; !ok {
				fmt.Println("Wrong!")
				continue
			}
			if median.Remove(v) {
				delete(dict, v)
				if median.lower.Size()+median.upper.Size() == 0 {
					fmt.Println("Wrong!")
				} else {
					fmt.Println(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", median.Get()), "0"), "."))
				}
			} else {
				fmt.Println("Wrong!")
			}
		} else {
			median.Add(v)
			dict[v] = 1
			fmt.Println(strings.TrimRight(strings.TrimRight(fmt.Sprintf("%f", median.Get()), "0"), "."))
		}
	}
}
