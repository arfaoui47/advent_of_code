package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	file, err := os.Open("2021/day9/input.txt")
	if err != nil {
		panic(err)
	}
	grid := [1000][1000]int{}
	scanner := bufio.NewScanner(file)
	line_count := 0
	line_lenght := 0
	for scanner.Scan() {
		line := scanner.Text()
		line_lenght = len(line)
		for index, char := range line {
			grid[line_count][index] = int(char) - 48
		}
		line_count += 1
	}
	sum := 0
	for i := 0; i < line_count; i++ {
		for j := 0; j < line_lenght; j++ {
			if less_than_adjacent(grid, i, j, line_lenght, line_count) {
				sum += grid[i][j] + 1
			}
		}
	}
	fmt.Println(sum)
}

func less_than_adjacent(grid [1000][1000]int, x int, y int, width int, height int) bool {
	left := false
	right := false
	up := false
	down := false
	if y < width-1 {
		if grid[x][y] < grid[x][y+1] {
			right = true
		}
	} else {
		right = true
	}
	if y > 0 {
		if grid[x][y] < grid[x][y-1] {
			left = true
		}
	} else {
		left = true
	}
	if x < height-1 {
		if grid[x][y] < grid[x+1][y] {
			up = true
		}
	} else {
		up = true
	}
	if x > 0 {
		if grid[x][y] < grid[x-1][y] {
			down = true
		}
	} else {
		down = true
	}
	return left && right && up && down
}

func main() {
	part1()
}
