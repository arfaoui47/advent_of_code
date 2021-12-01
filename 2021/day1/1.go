package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("2021/day1/input.txt")
	if err != nil {
		panic(err)
	}
	var nums []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	prev := nums[0]
	inremental := 0
	for _, num := range nums[1:] {
		if num > prev {
			inremental += 1
		}
		prev = num
	}
	println(inremental)
}
