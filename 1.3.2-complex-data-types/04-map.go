package main

import "fmt"

func main() {
	// Создание
	// [string] - тип ключа
	// string - тип значения

	var user map[string]string = map[string]string{
		"name":     "John",
		"lastName": "Doe",
	}

	// Задаем область памяти, так же как и со Slice - 10
	profile := make(map[string]string, 10)
	fmt.Println(profile)

	// Количество элементов
	mapLength := len(user)
	fmt.Println(mapLength)

	// Если ключа нет - вернет значеие по умолчанию для типа
	mName := user["middleName"]
	fmt.Println("mName", mName)

	// Проверка на существование ключа
	// При объявлении переменной пишем вторым параметром
	// Переменную где будет храниться, есть ли нет: true/false
	// Например:
	mName, mNameExist := user["middleName"]
	fmt.Println("mName:", mName, "mNameExist:", mNameExist)

	// Пустая переменная - только проверяем, что есть ключ
	_, mNmNameExist2 := user["middleName"]
	fmt.Println("mNmNameExist2:", mNmNameExist2)

	// Удаление ключа
	delete(user, "lastName")
	fmt.Println("user after delete lastName:", user)
}
