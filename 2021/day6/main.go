package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
	file, err := os.Open("2021/day6/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	fishes := []int{}
	fishes_st := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		fishes_st = strings.Split(line, ",")
	}
	fmt.Println(fishes_st, len(fishes_st))

	for i := 0; i < len(fishes_st); i++ {
		item, _ := strconv.Atoi(fishes_st[i])
		fishes = append(fishes, item)
	}
	for j := 0; j <= 80; j++ {
		fmt.Println("day ", j, len(fishes))
		for i := 0; i < len(fishes); i++ {

			if fishes[i] > 0 {
				fishes[i]--
			} else {
				fishes[i] = 6
				fishes = append(fishes, 9)
			}
		}
	}

}

func main() {
	part1()
}
