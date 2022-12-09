package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	ROCK_SCORE := 1
	PAPER_SCORE := 2
	SCISSORS_SCORE := 3

	LOSE_SCORE := 0
	DRAW_SCORE := 3
	WIN_SCORE := 6

	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)

	total_score := 0

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			split_line := strings.Split(line_str, " ")

			/* Syntax Notes
			 * Opponent Moves (split_line[0]):
			 * A == Rock
			 * B == Paper
			 * C == Scissors
			 *
			 * Your Moves (split_line[1])
			 * X == Lose
			 * Y == Draw
			 * Z == Win
			 */

			if split_line[0] == "A" && split_line[1] == "X" {
				total_score += SCISSORS_SCORE + LOSE_SCORE
			} else if split_line[0] == "A" && split_line[1] == "Y" {
				total_score += ROCK_SCORE + DRAW_SCORE
			} else if split_line[0] == "A" && split_line[1] == "Z" {
				total_score += PAPER_SCORE + WIN_SCORE
			} else if split_line[0] == "B" && split_line[1] == "X" {
				total_score += ROCK_SCORE + LOSE_SCORE
			} else if split_line[0] == "B" && split_line[1] == "Y" {
				total_score += PAPER_SCORE + DRAW_SCORE
			} else if split_line[0] == "B" && split_line[1] == "Z" {
				total_score += SCISSORS_SCORE + WIN_SCORE
			} else if split_line[0] == "C" && split_line[1] == "X" {
				total_score += PAPER_SCORE + LOSE_SCORE
			} else if split_line[0] == "C" && split_line[1] == "Y" {
				total_score += SCISSORS_SCORE + DRAW_SCORE
			} else if split_line[0] == "C" && split_line[1] == "Z" {
				total_score += ROCK_SCORE + WIN_SCORE
			}

		}

	}

	fmt.Printf("Total Score: %d\n", total_score)

}
