package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var cmp = func(i1 Item, i2 Item) int {
	i := i1.priority - i2.priority
	if i == 0 {
		return 0
	} else {
		return int(i / math.Abs(i))
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
		log.Fatal(err)
	}
	input = input[1:]
	fmt.Println(N)
	fmt.Println(input)
	fmt.Println("vim-go")
	pq := NewPQ(cmp)
	fmt.Println(pq.isEmpty())
}
