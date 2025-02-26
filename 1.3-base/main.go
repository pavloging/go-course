package main

import "fmt"

func main() {
	// Значние по умолчанию (Для int = 0)
	var num0 int

	// Значение при инциализации
	var num1 int = 1

	// Пропуск типа (Тип присваивается компилятором)
	var num2 = 20
	fmt.Println(num0, num1, num2)

	// Короткое объявление типа
	num := 30
	fmt.Println(num)
}
