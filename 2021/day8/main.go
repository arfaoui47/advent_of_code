package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func part1() {
	file, err := os.Open("2021/day8/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	sum_int := 0
	for scanner.Scan() {
		line := scanner.Text()
		reads := strings.Split(line, " | ")
		vals := strings.Split(reads[1], " ")
		for _, val := range vals {
			if len(val) == 2 || len(val) == 4 || len(val) == 3 || len(val) == 7 {
				sum_int += 1
			}
		}
	}
	fmt.Println(sum_int)
}

func main() {
	part1()
}
