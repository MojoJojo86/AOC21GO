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

	// PART 1
	measurements := 0

	for i := 1; i < len(values); i++ {
		if values[i] > values[i-1] {
			measurements++
		}
	}

	fmt.Print(measurements)
}
