package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Пустая строка по умолчанию ""
	var srt string
	fmt.Println(srt)

	// В двойныхковычках отображаются все передаваемые символы
	// Спец символы не применяются
	var hello string = "Привет/n/t"
	fmt.Println(hello)

	// В строке с бэк-слеш применяются спец символы
	var format string = `Мир\n\t`
	fmt.Println(format)

	// UTF-8 из коробки
	var helloWorld = "Привет, Мир!"
	hi := "你好, 世界"
	fmt.Println(helloWorld, hi)

	// Одинарные ковычки для байт (uint8)
	var rawBinary byte = '\x27'
	fmt.Println(rawBinary)

	// rune (unit32) для UTF-8 символов
	var someChinese rune = '你'
	fmt.Println("Код:", someChinese)

	// Конкотинация строк
	andGoodMorning := helloWorld + " и доброе утро!"
	fmt.Println(andGoodMorning)

	// Символ по индексу не изменяемый
	// helloWorld[0] = 72 - Ошибка cannot assign to helloWorld[0]

	// Получение длины строки
	byteLen := len(helloWorld)                    // 21 байт
	symbols := utf8.RuneCountInString(helloWorld) // 12 рун - 12 символов
	fmt.Println(byteLen, symbols)

	// Получение под строки в байтах
	hello = helloWorld[:12] // Обрезание в битах - Получим "Привет"
	H := helloWorld[0]      // Возвращает первый байт строки - Получим 208
	fmt.Println(hello, H)

	// Конвертация в slice байт и обратно
	byteString := []byte(helloWorld)
	helloWorld = string(byteString)
	fmt.Println(byteString, helloWorld)
}
