package main

import (
	"bufio"
	"fmt"
	"os"
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
func main() {
	part1()
}
