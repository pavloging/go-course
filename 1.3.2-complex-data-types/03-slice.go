package main

import "fmt"

func main() {
	buf := []int{1, 2, 3, 4, 5}
	fmt.Println(buf)

	// Получение среза, указывающего на ту же память
	// Slice - переводиться как срез

	sl1 := buf[1:4] // [2, 3, 4]
	sl2 := buf[:2]  // [1, 2]
	sl3 := buf[:2]  // [3, 4, 5]
	fmt.Println(sl1, sl2, sl3)

	newBuf := buf[:] // [1, 2, 3, 4, 5]
	newBuf[0] = 9    // Изменяем один из элементов
	// Сейчас buf обновился так как одна область памяти
	// buf = [9, 2, 3, 4, 5]

	// newBuf теперь указан на другие данные
	// Так как мы разширяем cap, то создается новая область памяти
	newBuf = append(newBuf, 6)

	// buf		= [9, 2, 3, 4, 5], не изменился
	// newBuf	= [1, 2, 3, 4, 5, 6], изменился
	newBuf[0] = 1
	fmt.Println("buf", buf)
	fmt.Println("newBuf", newBuf)

	// Копирование одного Slice в другой
	var emptyBuf []int // len=0, cap=0
	// Неправильно - скопирует меньшее (по len) из 2-х слайсов
	copied := copy(emptyBuf, buf) // copied = 0
	fmt.Println(copied, emptyBuf) // copied = 0

	// Правильно
	newBuf = make([]int, len(buf), len(buf))
	copy(newBuf, buf)
	fmt.Println(newBuf) // [9, 2, 3, 4, 5]

	// Можно копировать в часть существующего слайса
	ints := []int{1, 2, 3, 4}
	copy(ints[1:3], []int{5, 6})
	fmt.Println(ints) // [1, 5, 6, 4]
}
