package main

import (
	"fmt"
)

func main() {
	var array1 [20]float64
	var array2 [20]float64

	for i := 0; i < len(array1); i++ {
		array1[i] = 2.0
		array2[i] = 3.0
	}

	var suma1 float64 = 0.0
	var suma2 float64 = 0.0
	var suma3 float64 = 0.0

	for i := 0; i < len(array1); i++ {
		suma1 += array1[i]
		suma2 += array2[i]
		suma3 += array1[i] + array2[i]
	}

	fmt.Printf("Suma elementów tablicy 1: %f\n", suma1)
	fmt.Printf("Suma elementów tablicy 2: %f\n", suma2)
	fmt.Printf("Suma elementów tablic 1 i 2: %f\n", suma3)
}
