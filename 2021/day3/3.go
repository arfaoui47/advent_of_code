package main

import (
	"bufio"
	"os"
	"strconv"
)

func part1() {
	file, err := os.Open("2021/day3/input.txt")
	if err != nil {
		panic(err)
	}
	total := 1
	// store 1 counts
	m := make(map[int]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		total++
		for index, char := range line {
			if char == '1' {
				m[index]++
			}
		}
	}
	gamma := ""
	epsilon := ""
	for i := 0; i < len(m); i++ {
		if m[i] > total/2 {
			gamma += "1"
			epsilon += "0"
		} else {
			gamma += "0"
			epsilon += "1"
		}
	}
	println(gamma)
	println(epsilon)
	gamma_decimal, _ := binaryToDecimal(gamma)
	epsilon_decimal, _ := binaryToDecimal(epsilon)
	print(gamma_decimal * epsilon_decimal)

}

func binaryToDecimal(binary string) (int, error) {
	decimal, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		return 0, err
	}
	return int(decimal), nil
}

func main() {
	part1()
}
