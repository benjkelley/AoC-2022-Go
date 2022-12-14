package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Stack implementation from https://www.educative.io/answers/how-to-implement-a-stack-in-golang
type Stack []string

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *Stack) Print() {
	fmt.Println(*s)
}

func (s *Stack) Reverse() {
	var holder []string
	for s.IsEmpty() == false {
		temp, suc := s.Pop()
		if suc == true {
			holder = append(holder, temp)
		}
	}
	for i := 0; i < len(holder); i++ {
		s.Push(holder[i])
	}
}

func main() {
	inputs, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(inputs)
	diagram_counter := 0
	var stack_holder [9]Stack

	for scanner.Scan() {
		line_str := scanner.Text()
		if line_str != "" {
			if diagram_counter < 8 {
				split_line := strings.Split(line_str, "")
				position := 1
				for i := 0; i < 9; i++ {
					if split_line[position] != " " {
						stack_holder[i].Push(split_line[position])
					}
					position += 4
				}
			} else if diagram_counter == 8 {
				for i := 0; i < 9; i++ {
					stack_holder[i].Reverse()
				}
			} else if diagram_counter >= 9 {
				split_line := strings.Split(line_str, " ")
				//fmt.Println(split_line)
				count, _ := strconv.Atoi(split_line[1])
				from_col, _ := strconv.Atoi(split_line[3])
				to_col, _ := strconv.Atoi(split_line[5])

				var holder []string

				for i := 0; i < count; i++ {
					//fmt.Printf("Loop %d, count %d", i, count)
					temp, suc := stack_holder[from_col-1].Pop()
					if suc == false {
						fmt.Printf("ERROR: Stack %d Empty", from_col)
					}
					holder = append(holder, temp)
				}

				for i := len(holder) - 1; i >= 0; i-- {
					stack_holder[to_col-1].Push(holder[i])
				}

			}
			diagram_counter++
		}

	}

	answer := ""
	for i, x := range stack_holder {
		fmt.Println(x)
		reply, suc := x.Pop()

		if suc != false {
			answer = answer + reply
		} else {
			fmt.Printf("Stack %d is empty\n", i+1)
		}
	}

	fmt.Printf("Final tops: %s\n", answer)

}
