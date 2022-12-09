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

	scanner := bufio.NewScanner(inputs)
	diagram_counter := 0

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			if diagram_counter < 9 {
				split_line := strings.Split(line_str, "")
				fmt.Println(split_line)
				diagram_counter++
			}
		}
	}
}
