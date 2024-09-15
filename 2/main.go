package main

import (
	"fmt"
	"os"
)

func main() {
	// Реализовать функцию, которая принимает число и возвращает "Positive", "Negative" или "Zero".

	// создание переменной
	var x int

	// получение значения переменной
	fmt.Print("введите число: ")
	fmt.Fscan(os.Stdin, &x)

	// определение знака числа
	if x > 0 {
		fmt.Println("Positive")
	} else if x < 0 {
		fmt.Println("Negative")
	} else {
		fmt.Println("Zero")
	}

}
