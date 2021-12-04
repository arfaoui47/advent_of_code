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
	gamma_decimal, _ := binaryToDecimal(gamma)
	epsilon_decimal, _ := binaryToDecimal(epsilon)
	println(gamma_decimal * epsilon_decimal)

}

func part2() {
	file, err := os.Open("2021/day3/input.txt")
	if err != nil {
		panic(err)
	}
	total := 1
	// store 1 counts
	m := make(map[int]int)
	list := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		total++
		list = append(list, line)
		for index, char := range line {
			if char == '1' {
				m[index]++
			}
		}
	}
	co2_list := list
	oxygen := make([]string, 0)
	aux_m_oxy := m

	for i := 0; i < len(aux_m_oxy); i++ {
		oxygen = []string{}

		if aux_m_oxy[i] >= (len(list)+1)/2 {
			for _, v := range list {
				if v[i] == '1' {
					oxygen = append(oxygen, v)
				}
			}
		} else {
			for _, v := range list {
				if v[i] == '0' {
					oxygen = append(oxygen, v)
				}
			}
		}
		list = oxygen
		aux_m_oxy = make(map[int]int)
		for _, line := range list {
			for index, char := range line {
				if char == '1' {
					aux_m_oxy[index] += 1
				} else {
					aux_m_oxy[index] += 0
				}
			}
		}
	}

	aux_m_co2 := m
	co2 := make([]string, 0)
	for i := 0; i < len(aux_m_co2); i++ {
		co2 = []string{}

		if aux_m_co2[i] < (len(co2_list)+1)/2 {
			for _, v := range co2_list {
				if v[i] == '1' {
					co2 = append(co2, v)
				}
			}
		} else {
			for _, v := range co2_list {
				if v[i] == '0' {
					co2 = append(co2, v)
				}
			}
		}
		co2_list = co2
		aux_m_co2 = make(map[int]int)
		for _, line := range co2_list {
			for index, char := range line {
				if char == '1' {
					aux_m_co2[index] += 1
				} else {
					aux_m_co2[index] += 0
				}
			}
		}
		if len(co2) == 1 {
			break
		}
	}
	oxy := oxygen[0]
	co2_ := co2[0]

	oxy_decimal, _ := binaryToDecimal(oxy)
	co2_decimal, _ := binaryToDecimal(co2_)
	println("part2", co2_decimal*oxy_decimal)
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
	part2()
}
