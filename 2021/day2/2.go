package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("2021/day2/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	horizontal := 0
	vertical := 0
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, " ")
		direction := v[0]
		value, _ := strconv.Atoi(v[1])
		if direction == "forward" {
			horizontal += value
		} else if direction == "down" {
			vertical += value
		} else if direction == "up" {
			vertical -= value
		}
	}
	println(horizontal * vertical)
}

func part2() {
	file, err := os.Open("2021/day2/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	horizontal := 0
	vertical := 0
	aim := 0
	for scanner.Scan() {
		line := scanner.Text()
		v := strings.Split(line, " ")
		direction := v[0]
		value, _ := strconv.Atoi(v[1])
		if direction == "forward" {
			horizontal += value
			vertical += aim * value
		} else if direction == "down" {
			aim += value
		} else if direction == "up" {
			aim -= value
		}
	}
	println(horizontal * vertical)
}

func main() {
	part1()
	part2()
}
