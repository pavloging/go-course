# Переменные

Создание перменных используются ключевое слово var [имя] [тип] = [значение], например:

```go
var car string = "BMW"
```

Какие типы данных есть в Go?
TODO: Дописать или отредактировать

1. Простые типы:

    - Целочисленные типы:
        - int
        - int8
        - int16
        - int32
        - int64
        - uint
        - uint8 (байт)
        - uint16
        - uint32
        - uint64
        - uintptr
    - Числа с плавающей запятой:
        - float32
        - float64
    - Комплексные числа:
        - complex64
        - complex128
    - Булевый тип:
        - bool
    - Строки:
        - string

2. Составные типы:

    - Массивы: фиксированная длина, например, [5]int — массив из 5 целых чисел.
    - Срезы: динамические массивы, например, []int — срез целых чисел.
    - Словари (карты): неупорядоченные коллекции пар ключ-значение, например, map[string]int — карта, где ключи — строки, а значения — целые числа.
    - Структуры: пользовательские типы, которые могут содержать несколько полей, например:

    ```go
         type Person struct {
          Name string
          Age  int
      }
    ```

3. Интерфейсы: позволяют определять методы, которые должны реализовывать типы, например:

```go
   type Reader interface {
       Read(p []byte) (n int, err error)
   }
```

4. Функции: функции могут быть типами, например:

```go
   type IntFunc func(int) int
```

5. Указатели: типы, которые хранят адреса других переменных, например, \*int — указатель на целое число.

6. Типы каналов: используются для синхронизации между горутинами, например:

```go
   ch := make(chan int)
```
