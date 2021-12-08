package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("2021/day7/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	fuels := []int{}
	fuels_st := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fuels_st = strings.Split(line, ",")
	}
	for i := 0; i < len(fuels_st); i++ {
		fuel, _ := strconv.Atoi(fuels_st[i])
		fuels = append(fuels, fuel)
	}
	min_fuels := int(math.Pow(10, 10))
	sort.Ints(fuels)
	min_ := fuels[0]
	max_ := fuels[len(fuels)-1]
	for i := min_; i <= max_; i++ {
		sum := 0
		for j := 0; j < len(fuels); j++ {
			sum += Abs(fuels[j] - i)
		}
		if sum < min_fuels {
			min_fuels = sum
		}
	}
	fmt.Println(min_fuels)
}

func part2() {
	file, err := os.Open("2021/day7/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	fuels := []int{}
	fuels_st := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fuels_st = strings.Split(line, ",")
	}
	for i := 0; i < len(fuels_st); i++ {
		fuel, _ := strconv.Atoi(fuels_st[i])
		fuels = append(fuels, fuel)
	}
	min_fuels := int(math.Pow(10, 10))
	sort.Ints(fuels)
	min_ := fuels[0]
	max_ := fuels[len(fuels)-1]
	for i := min_; i <= max_; i++ {
		sum := 0
		for j := 0; j < len(fuels); j++ {
			for k := 0; k < Abs(fuels[j]-i)+1; k++ {
				sum += k
			}
		}
		if sum < min_fuels {
			min_fuels = sum
		}
	}
	fmt.Println(min_fuels)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	part1()
	part2()
}
