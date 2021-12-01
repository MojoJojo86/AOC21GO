package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer data.Close()

	scanner := bufio.NewScanner(data)
	scanner.Split(bufio.ScanLines)

	var values []int
	for scanner.Scan() {

		num, _ := strconv.Atoi(scanner.Text())
		values = append(values, num)
	}

	// PART 2
	measurements := 0

	// i+3 to check there are enough measurements
	for i := 0; i+3< len(values); i++ {
		window1 := values[i] + values[i+1] + values[i+2]
		window2 := values[i+1] + values[i+2] + values[i+3]

		if window2 > window1 {
			measurements++
		}
	}

	fmt.Print(measurements)
}
