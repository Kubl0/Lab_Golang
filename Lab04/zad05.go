package main

import "fmt"

func main() {
	var matrix1 = make([][]int, 3)
	var matrix2 = make([][]int, 3)

	for i := 0; i < 3; i++ {
		matrix1[i] = make([]int, 3)
		matrix2[i] = make([]int, 3)

		for j := 0; j < 3; j++ {
			matrix1[i][j] = i + j
			matrix2[i][j] = (i + 1) * j
		}
	}

	fmt.Println("Matrix 1:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix1[i][j])
		}
		fmt.Println()
	}
	fmt.Println("Matrix 2:")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix2[i][j])
		}
		fmt.Println()
	}

}
