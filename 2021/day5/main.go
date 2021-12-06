package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	point struct {
		x int
		y int
	}
	coord struct {
		start point
		end   point
	}
	coords []coord
)

func part1() {
	file, err := os.Open("2021/day5/input.txt")
	if err != nil {
		panic(err)
	}
	// store 1 counts
	coordinates := make(coords, 0)
	scanner := bufio.NewScanner(file)
	maxX := 0
	maxY := 0

	for scanner.Scan() {
		line := scanner.Text()
		cords := strings.Split(line, " -> ")
		start := strings.Split(cords[0], ",")
		startX, _ := strconv.Atoi(start[1])
		startY, _ := strconv.Atoi(start[0])
		cord_start := point{startX, startY}
		end := strings.Split(cords[1], ",")
		endX, _ := strconv.Atoi(end[1])
		endY, _ := strconv.Atoi(end[0])
		cord_end := point{endX, endY}
		coordinates = append(coordinates, coord{cord_start, cord_end})
	}
	grid := [1000][1000]int{}
	for i := 0; i < maxX; i++ {
		for j := 0; j < maxY; j++ {
			grid[i][j] = 0
		}
	}

	for _, cord := range coordinates {
		if (cord.start.x != cord.end.x) && (cord.start.y != cord.end.y) {
			continue
		}
		if cord.start.x == cord.end.x {
			if cord.start.y < cord.end.y {
				for j := cord.start.y; j <= cord.end.y; j++ {
					grid[cord.end.x][j] += 1
				}
			} else {
				for j := cord.end.y; j <= cord.start.y; j++ {
					grid[cord.end.x][j] += 1
				}
			}
		}
		if cord.start.y == cord.end.y {
			if cord.start.x < cord.end.x {
				for i := cord.start.x; i <= cord.end.x; i++ {
					grid[i][cord.end.y] += 1
				}
			} else {
				for i := cord.end.x; i <= cord.start.x; i++ {
					grid[i][cord.end.y] += 1
				}
			}
		}
	}
	sum := 0
	for _, row := range grid {
		for _, v := range row {
			if v > 1 {
				sum++
			}
		}
	}
	fmt.Println(sum)
}

func main() {
	part1()
}
