package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func part1() {
	file, err := os.Open("2021/day11/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	levels := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		j_level := []int{}
		for _, c := range line {
			level, _ := strconv.Atoi(string(c))
			j_level = append(j_level, level)
		}
		levels = append(levels, j_level)
	}
	for turn := 0; turn < 2; turn++ {
		for i := 0; i < len(levels); i++ {
			for j := 0; j < len(levels[i]); j++ {
				// increment level
				if levels[i][j] == 9 {
					levels[i][j] = 0
				} else if levels[i][j] != 0 {
					levels[i][j]++
				}
				// check if it will glow

			}
		}
		for i := 0; i < len(levels); i++ {
			for j := 0; j < len(levels[i]); j++ {
				if levels[i][j] == 0 {
					incrementNeighbours(levels, i, j)
				}
			}
		}
		for i := 0; i < len(levels); i++ {
			for j := 0; j < len(levels[i]); j++ {
				fmt.Print(levels[i][j])
			}
			fmt.Println()
		}
		fmt.Println()
	}

}

func incrementLevel(levels [][]int, i int, j int) {
	if levels[i][j] == 9 {
		levels[i][j] = 0
	} else if levels[i][j] != 0 {
		levels[i][j]++
	}
}

func incrementNeighbours(levels [][]int, i int, j int) {
	if i > 0 {
		if levels[i-1][j] != 0 {
			incrementLevel(levels, i-1, j)
		}
		if j > 0 {
			if levels[i-1][j-1] != 0 {
				incrementLevel(levels, i-1, j-1)
			}
		}
		if j < len(levels[i])-1 {
			if levels[i-1][j+1] != 0 {
				incrementLevel(levels, i-1, j+1)
			}
		}
	}
	if i < len(levels)-1 {
		if levels[i+1][j] != 0 {
			incrementLevel(levels, i+1, j)
		}
		if j > 0 {
			if levels[i+1][j-1] != 0 {
				incrementLevel(levels, i+1, j-1)
			}
		}
		if j < len(levels[i])-1 {
			if levels[i+1][j+1] != 0 {
				incrementLevel(levels, i+1, j+1)
			}
		}
	}
	if j > 0 {
		if levels[i][j-1] != 0 {
			incrementLevel(levels, i, j-1)
		}
	}
	if j < len(levels[i])-1 {
		if levels[i][j+1] != 0 {
			incrementLevel(levels, i, j+1)
		}
	}
}

func main() {
	part1()
}
