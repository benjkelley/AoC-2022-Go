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
	highest_calories := 0

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			num, _ := strconv.Atoi(line_str)
			calories_counter += num
		} else {
			if calories_counter > highest_calories {
				highest_calories = calories_counter
			}
			calories_counter = 0
		}
	}

	fmt.Printf("Highest Calories Count: %d\n", highest_calories)

}
