package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)

	calories_counter := 0
	highest_calories := [3]int{0, 0, 0}

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			num, _ := strconv.Atoi(line_str)
			calories_counter += num
		} else {
			if calories_counter > highest_calories[1] {
				if calories_counter > highest_calories[2] {
					highest_calories[0] = highest_calories[1]
					highest_calories[1] = highest_calories[2]
					highest_calories[2] = calories_counter
				} else {
					highest_calories[0] = highest_calories[1]
					highest_calories[1] = calories_counter
				}
			} else if calories_counter < highest_calories[1] {
				if calories_counter > highest_calories[0] {
					highest_calories[0] = calories_counter
				}
			}
			calories_counter = 0
		}

	}

	calories_sum := 0
	for _, n := range highest_calories {
		calories_sum += n
	}

	fmt.Printf("Highest Calories Count: %d\n", calories_sum)
}
