package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)

	duplicate_counter := 0

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			split_line := strings.Split(line_str, ",")
			elf_1 := strings.Split(split_line[0], "-")
			elf_2 := strings.Split(split_line[1], "-")

			elf_1_low, _ := strconv.Atoi(elf_1[0])
			elf_1_high, _ := strconv.Atoi(elf_1[1])
			elf_2_low, _ := strconv.Atoi(elf_2[0])
			elf_2_high, _ := strconv.Atoi(elf_2[1])

			if elf_1_low <= elf_2_low && elf_1_high >= elf_2_low { // 1 overlaps with 2 on lower range
				duplicate_counter++
			} else if elf_1_low <= elf_2_high && elf_1_high >= elf_2_high { // 1 overlaps 2 on high range
				duplicate_counter++
			} else if elf_2_low <= elf_1_low && elf_2_high >= elf_1_low { // 2 overlaps with 1 on lower range
				duplicate_counter++
			} else if elf_2_low <= elf_1_high && elf_2_high >= elf_1_high { // 2 overlaps 1 on high range
				duplicate_counter++
			}
		}
	}

	fmt.Printf("Duplicate Ranges: %d\n", duplicate_counter)
}
