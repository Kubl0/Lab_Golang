package main

import (
	"fmt"
)

func main() {
	//test_wynikow()
	silne_liczby()
}

func test_wynikow() {
	var wynik int
	wynik = 6 / 2 * (1 + 2)
	fmt.Println("Wynik 1:", wynik)
	var wynik2 float64
	wynik2 = 9 - 3*3 + 1
	fmt.Println("Wynik 2:", wynik2)
}

func silnia(value int) int {
	if value == 0 {
		return 1
	}

	return value * silnia(value-1)
}

func silne_liczby() {
	var imie string
	var nazwisko string

	fmt.Println("Podaj imie")
	fmt.Scanln(&imie)
	fmt.Println("Podaj nazwisko")
	fmt.Scanln(&nazwisko)

	imie = imie[0:3]
	nazwisko = nazwisko[0:3]

	nick := imie + nazwisko

	fmt.Println(nick)

	var ascii_array [6]int

	for i := 0; i < len(nick); i++ {
		ascii_array[i] = int(nick[i])
	}

	fmt.Println(ascii_array)

	var found bool = false
	j := 1

	for found != true {
		fmt.Println(silnia(j))
		j++
		if j > 10 {
			found = true
		}
	}
}
