package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	total_priority := 0

	var values map[string]int
	values = make(map[string]int)

	all_letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for i, x := range all_letters {
		values[string(x)] = i + 1
	}

	scanner := bufio.NewScanner(inputs)

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			split_line := strings.Split(line_str, "")
			midpoint := len(split_line) / 2
			section_one := split_line[0:midpoint]
			section_two := split_line[midpoint:]

			var duplicate_char string

			for _, x := range section_one {
				for _, y := range section_two {
					if x == y {
						duplicate_char = x
						break
					}
				}
				if duplicate_char != "" {
					break // there should only be one duplicate
				}
			}

			if duplicate_char == "" {
				fmt.Println(section_one)
				fmt.Println(section_two)
			} else {
				fmt.Printf("Adding value %d for character %s\n", values[string(duplicate_char)], duplicate_char)
				total_priority += values[string(duplicate_char)]
			}

		}
	}

	fmt.Printf("Total Priority: %d\n", total_priority)

}
