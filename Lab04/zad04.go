package main

import "fmt"

func main() {
	var size = 10

	var slice1 = make([][]int, size)

	var slice2 = make([][]int, size)

	var numbers [100]int

	for i := 0; i < 100; i++ {
		numbers[i] = i
	}

	for i := 0; i < size; i++ {
		slice1[i] = numbers[i*size : i*size+size]
	}

	for i := 0; i < size; i++ {
		slice2[i] = make([]int, size)
		for j := 0; j < size; j++ {
			slice2[i][j] = slice1[size-1-i][size-1-j]
		}
	}

	var resultMatrix = make([][]int, size)

	for i := 0; i < size; i++ {
		resultMatrix[i] = make([]int, size)
		for j := 0; j < size; j++ {
			resultMatrix[i][j] = slice1[i][j] + slice2[i][j]
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			fmt.Printf("%d ", resultMatrix[i][j])
		}
		fmt.Println()
	}

}
