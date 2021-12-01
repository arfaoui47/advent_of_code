package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1() {
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

func part2() {
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

	prev := nums[0] + nums[1] + nums[2]
	inremental := 0

	for i := 1; i < len(nums)-2; i += 1 {
		current := nums[i] + nums[i+1] + nums[i+2]
		if current > prev {
			inremental += 1
		}
		prev = current
	}
	println(inremental)
}

func main() {
	part1()
	part2()
}
