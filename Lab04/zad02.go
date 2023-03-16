package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var macierz1 [3][3]int
	var macierz2 [3][3]int

	for i := 0; i < len(macierz1); i++ {
		for j := 0; j < len(macierz1); j++ {
			macierz1[i][j] = rand.Intn(10)
			macierz2[i][j] = rand.Intn(10)
		}
	}

	var resultMatrix [3][3]int

	for i := 0; i < len(macierz1); i++ {
		for j := 0; j < len(macierz1); j++ {
			resultMatrix[i][j] = macierz1[i][j] * macierz2[i][j]
		}
	}

	for i := 0; i < len(macierz1); i++ {
		fmt.Printf("%d%s%d", macierz1[i], "  ", macierz2[i])
		fmt.Println()
	}
	fmt.Println("Macierz wynikowa:", resultMatrix)

}
