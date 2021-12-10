package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func part1() {
	file, err := os.Open("2021/day10/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	stack := make([]string, 0)
	resut := 0
	mapping := map[string]string{"{": "}", "[": "]", "(": ")", "<": ">"}
	for scanner.Scan() {
		line := scanner.Text()
		for _, char := range line {
			if char == '{' || char == '[' || char == '(' || char == '<' {
				stack = append(stack, string(char))
			} else if char == '}' || char == ']' || char == ')' || char == '>' {
				if len(stack) == 0 {
					panic("stack is empty")
				}
				if mapping[stack[len(stack)-1]] != string(char) {
					if char == '}' {
						resut += 1197
					} else if char == ']' {
						resut += 57
					} else if char == ')' {
						resut += 3
					} else if char == '>' {
						resut += 25137
					}

				}
				stack = stack[:len(stack)-1]
			}
		}

	}
	fmt.Println(resut)
}
func part2() {
	file, err := os.Open("2021/day10/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	corrupted := false
	total_scores := make([]int, 0)
	stack := make([]string, 0)
	mapping := map[string]string{"{": "}", "[": "]", "(": ")", "<": ">"}
	for scanner.Scan() {
		line := scanner.Text()
		corrupted = false
		stack = make([]string, 0)
		for _, char := range line {
			if char == '{' || char == '[' || char == '(' || char == '<' {
				stack = append(stack, string(char))
			} else if char == '}' || char == ']' || char == ')' || char == '>' {
				if len(stack) == 0 {
					panic("stack is empty")
				}
				if mapping[stack[len(stack)-1]] != string(char) {
					corrupted = true
					break
				}
				stack = stack[:len(stack)-1]
			}
		}
		score := 0
		var v int
		if !corrupted {
			for i := len(stack) - 1; i >= 0; i-- {
				if stack[i] == "(" {
					v = 1
				} else if stack[i] == "[" {
					v = 2
				} else if stack[i] == "{" {
					v = 3
				} else if stack[i] == "<" {
					v = 4
				}
				score = 5*score + v
			}
			total_scores = append(total_scores, score)
		}

	}

	sort.Ints(total_scores)
	fmt.Println(total_scores[len(total_scores)/2])

}
func main() {
	// part1()
	part2()
}
