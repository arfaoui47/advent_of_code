package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Maps map[int]int

	Board struct {
		values      [5][5]string
		row_match   Maps
		colum_match Maps
	}
)

func part1() {
	file, err := os.Open("2021/day4/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	draws := strings.Split(lines[0], ",")
	boards := []Board{}
	for i := 2; i < len(lines); i += 6 {
		b := Board{}
		b.colum_match = make(Maps)
		b.row_match = make(Maps)
		for j := 0; j < 5; j++ {
			row := strings.Fields(lines[i+j])
			for k := 0; k < 5; k++ {
				b.values[j][k] = row[k]
			}
		}
		boards = append(boards, b)
	}
	winner_values := 0
	winner_draw := ""

	drawed := []string{}

	for _, draw := range draws {
		drawed = append(drawed, draw)
		for _, board := range boards {
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if board.values[i][j] == draw {
						board.row_match[i]++
						board.colum_match[j]++
					}
					if board.row_match[i] == 5 || board.colum_match[j] == 5 {

						winner_draw = draw
						for k := 0; k < 5; k++ {
							for m := 0; m < 5; m++ {
								if !contains(drawed, board.values[k][m]) {
									value, _ := strconv.Atoi(board.values[k][m])
									winner_values += value
								}
							}
						}
						break

					}
				}

				if winner_draw != "" {
					break
				}

			}
			if winner_draw != "" {
				break
			}
		}
		if winner_draw != "" {
			break
		}
	}
	winner_draw_int, _ := strconv.Atoi(winner_draw)
	fmt.Println(winner_values * winner_draw_int)
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func main() {
	part1()
}
