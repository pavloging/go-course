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

	// Повторное объявление через num := 42 - не отработает!

	// Полная запись суммы
	num = num + 1
	fmt.Println("num = num + 1", num)

	// Краткая запись суммы
	num += 1
	fmt.Println("+=", num)

	// Более краткое добавление переменной еденицы
	num++
	fmt.Println("++", num)

	// Стиль объявление переменных - camelCase
	// Не принятый стиль - under_score
	userIndex := 10
	under_score := 10
	fmt.Println("Используем camelCase", userIndex)
	fmt.Println("Не используем under_score", under_score)

	// Объявление переменных
	var weight, height int = 76, 185

	// Присваивание существующим переменным новые значения
	weight, height = 80, 190

	// Короткое присваивание
	// Но хотя бы одна из переменных должна быть новой
	weight, age := 25, 90

	fmt.Println(weight, height, age)
}
