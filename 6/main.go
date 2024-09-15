package main

import (
	"fmt"
	"os"
)

func main() {
	// Написать функцию, которая принимает два целых числа и возвращает их среднее значение.

	// создание переменных
	var x float64
	var y float64

	// получение значения переменных
	fmt.Print("введите числа: ")
	fmt.Fscan(os.Stdin, &x)
	fmt.Fscan(os.Stdin, &y)

	// вызов функции
	average(x, y)

}

// функция для вычисления среднего значения двух чисел
func average(x float64, y float64) {
	a := (x + y) / 2
	fmt.Println("среднее значение: ", a)
}
