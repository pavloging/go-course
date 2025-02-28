package main

import "fmt"

func main() {

	// Значения по умолчанию
	var a1 [3]int // [0, 0, 0]
	fmt.Println("a1", a1)

	// Для определения размера массива используются только const
	const size = 2
	var a2 [2 * size]bool // [false,false,false]
	fmt.Println("a2", a2)

	// определение размера при объявлении
	a3 := [...]int{1, 2, 3}
	fmt.Println("a2", a3)

	// проверка при компиляции или при выполнении
	// invalid array index 4 (out of bounds for 3-element array)
	// a3[4] = 12
}
