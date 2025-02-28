package main

import "fmt"

const pi = 3.14
const (
	hello = "Привет"
	e     = 2.718
)
const (
	zero = iota // = 0
	_           // Пропуск переменной, но +1 будет -> = 1
	two         // = 2
)

const (
	_         = iota             // Пропускаем первое значение
	KB uint64 = 1 << (10 * iota) // 1024
	MB                           // 1048576
)

const (
	// Нетипизированная константа
	year = 2025
	// Типизированная констанат
	yearTyped int = 2025
)

func main() {
	var month int32 = 13
	fmt.Println(month + year)

	// fmt.Println(month + yearTyped) Ошибка - invalid operation: month + yearTyped (mismatched types int32 and int)
	fmt.Println(month + int32(yearTyped))
}
