package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func contains(str string, list []string) bool {
	for _, x := range list {
		if x == str {
			return true
		}
	}

	return false
}

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
	var groups_array [3][]string
	inner_loop_ctr := 0

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			split_line := strings.Split(line_str, "")
			groups_array[inner_loop_ctr] = split_line

			if inner_loop_ctr < 2 {
				inner_loop_ctr++

			} else {
				var first_two_shared []string
				var shared_item string
				for _, x := range groups_array[0] {
					if contains(x, groups_array[1]) == true && contains(x, first_two_shared) == false {
						first_two_shared = append(first_two_shared, string(x))
					}
				}

				fmt.Printf("Shared items for first two:\n")
				fmt.Println(first_two_shared)

				for _, x := range first_two_shared {
					if contains(x, groups_array[2]) {
						shared_item = x
						fmt.Printf("Found third shared item %s \n", shared_item)
						break
					}
				}
				total_priority += values[string(shared_item)]
				inner_loop_ctr = 0
			}
		}
	}

	fmt.Printf("Total Priority: %d\n", total_priority)

}
