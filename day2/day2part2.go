package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	horizontal := 0
	depth := 0
	move := 0
	aim := 0

	for i := 0; i < len(values); i++ {

		command := strings.Fields(values[i])
		move, _ = strconv.Atoi(command[1])

		if command[0] == "forward" {
			horizontal += move
			depth += aim * move
		} else if command[0] == "up" {
			aim -= move
		} else if command[0] == "down" {
			aim += move
		}
	}

	fmt.Print(depth * horizontal)

}
