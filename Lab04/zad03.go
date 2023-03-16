package main

import (
	"fmt"
)

func main() {
	//wycinek 3x3
	var slice1 = make([][]int, 3)

	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for i := 0; i < 3; i++ {
		slice1[i] = numbers[i*3 : i*3+3]
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", slice1[i][j])
		}
		fmt.Println()
	}
}
