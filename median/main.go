package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

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
