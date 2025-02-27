package main

import "fmt"

func main() {
	// int - целоцисленный тип + платформозависемый тип, либо 32 бита либо 64 бита (зависит от OC)
	var i int = 42

	// Тип выбирается автоматически если не передавать его
	var autoInt = -10

	// Можем настоить разрядность - int8, int16, int32, int64
	var bigInt int64 = 1<<32 - 1

	// uint - только для целых положительных чисел включая 0 тип + платформозависемый тип, либо 32 бита либо 64 бита (зависит от OC)
	var unsignedInt1 uint = 42
	var unsignedInt2 uint = 0
	// var unsignedInt3 uint = -42 - будет ошибка

	// Можем настоить разрядность - uint8, uint16, uint32, uint64
	var unsignedBigInt int64 = 1<<32 - 1

	fmt.Println(i, autoInt, bigInt, unsignedInt1, unsignedInt2, unsignedBigInt)

	var pi float32 = 3.141592653589
	var e float64 = 2.718281828459045
	goldenRatio := 1.1415926535

	fmt.Println(pi, e, goldenRatio)

	// bool - в качестве значения имеет true/false
	var b bool // false - по умолчанию
	var isOk bool = true
	var success = true
	fail := false

	fmt.Println(b, isOk, success, fail)

	// complex64, complex128
	// Что такое комплексные числа?
	// Комплексные переменные — это переменные, принимающие значения в виде комплексных чисел.
	// Комплексное число имеет форму ( z = a + bi ), где ( a ) и ( b ) — действительные числа,
	// ( a ) называется действительной частью, ( b ) — мнимой частью, а ( i ) — мнимая единица,
	// для которой выполняется равенство ( i^2 = -1 ).
	// Комплексные переменные часто используются в различных областях математики и физики.
	var c complex64 = 1 + 2i // 1 - действительная часть, 2 - мнимая часть
	d := 2 + 3i

	fmt.Println(c, d)
}
