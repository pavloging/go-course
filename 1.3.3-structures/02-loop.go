package main

import "fmt"

func main() {
	// Цикл без условия, while(true) of for(;;;)
	for {
		fmt.Println("loop iteration")
		break
	}

	// Цикл без условия, while(isRun)
	isRun := true
	for isRun {
		fmt.Println("loop iteration with condition")
		isRun = false
	}

	// Цикл с условием и блоком инициализации
	for i := 0; i < 2; i++ {
		fmt.Println("loop iteration", i)
		if i == 1 {
			continue // Следующая итерация цикла - continue
		}
	}

	// Операции по Slice
	sl := []int{1, 2, 3}
	idx := 0

	for idx < len(sl) {
		fmt.Println("while-style loop, idx:", idx, "value:", sl[idx])
		idx++
	}

	for i := 0; i < len(sl); i++ {
		fmt.Println("c-style loop", i, sl[i])
	}

	// Аналог for с индексом, но более упращенный
	for idx, val := range sl {
		fmt.Println("range loop", idx, val)
	}

	// Операции по map
	profile := map[int]string{1: "Vasily", 2: "Romanov"}

	// Перебор ключей (ключи не сохраняют порядок добавления поэтому
	// результат в map, может быть хаотичен
	// и значение будет не {"1": "1", "2": "2"}, наоборот)
	for key := range profile {
		fmt.Println("range map by key", key)
	}

	// Перебор ключей и значений
	for key, val := range profile {
		fmt.Println("range map by key-val", key, val)
	}

	// Перебор значений
	for _, val := range profile {
		fmt.Println("range map by val", val)
	}

	str := "Привет, Мир!"
	for pos, char := range str {
		// fmt.Printf("%#U at pos %dn", char, pos)
		fmt.Println(char, pos)
		// pos — это индекс текущего символа в строке
		// char — это текущий символ, представленный в виде руны (rune)

		if char == '!' {
			println("Win")
		}

	}

}
