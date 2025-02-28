package main

import "fmt"

func main() {
	// Простое условие
	boolVal := true
	if boolVal {
		fmt.Println("boolVal is true")
	}

	mapVal := map[string]string{"name": "rvasily"}
	// Условие с блоком инциализации
	if keyValue, keyExist := mapVal["name"]; keyExist {
		fmt.Println("name", keyValue)
	}

	// Получаем только признак сущестования ключа
	if _, keyExist := mapVal["name"]; keyExist {
		fmt.Println("key 'name' exist")
	}

	coud := 1
	// Блоки if/else
	if coud == 1 {
		fmt.Println("coud is 1")
	} else if coud == 2 {
		fmt.Println("coud is 2")
	} else {
		fmt.Println("coud is not 1 or 2")
	}

	// Блоки switch/case
	strVal := "name"
	switch strVal {
	case "name":
		fallthrough // Продолжить выполнение следующего случая
	case "test", "lastName":
		// some work
	default:
		// some work
	}

	// switch как замена многим if/else
	var val1, val2 = 2, 2
	switch {
	case val1 > 1 || val1 < 11:
		fmt.Println("val1 is less than val2")
	case val2 > 10:
		fmt.Println("val1 is greater than val2")
	default:
		fmt.Println("val1 is equal to val2")
	}
	// Выход из цикла, находясь внутри switch
Loop:
	for key, val := range mapVal {
		println("switch in loop", key, val)
		switch {
		case key == "lastName":
			break
			println("don`t pront this")
		case key == "firstName" && val == "Vasily":
			println("switch - break loop here")
			break Loop // Завершаем цикл с меткой "Loop:"
		}
	} // конец for

}
