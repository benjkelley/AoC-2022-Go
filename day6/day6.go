package main

import (
	"bufio"
	"fmt"
	"os"
)

func check_unique(char_holder []string) bool {
	for i, x := range char_holder {
		for j, y := range char_holder {
			if i != j {
				if x == y {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)

	for scanner.Scan() {
		line_str := scanner.Text()

		var char_holder []string

		for i, x := range line_str {
			if len(char_holder) < 3 {
				char_holder = append(char_holder, string(x))
			} else if len(char_holder) == 3 {
				char_holder = append(char_holder, string(x))

				if check_unique(char_holder) == true {
					fmt.Printf("Found non-repeating pattern at index %d\n", i+1)
					fmt.Println(char_holder)
					return
				}
			} else {
				char_holder[0] = char_holder[1]
				char_holder[1] = char_holder[2]
				char_holder[2] = char_holder[3]
				char_holder[3] = string(x)

				if check_unique(char_holder) == true {
					fmt.Printf("Found non-repeating pattern at index %d\n", i+1)
					fmt.Println(char_holder)
					return
				}
			}
		}
	}
	fmt.Println("No non-repeating pattern found")

}
