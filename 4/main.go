package main

import (
	"fmt"
	"os"
)

func main() {
	// Написать функцию, которая принимает строку и возвращает ее длину.

	// создание переменной
	var str string

	// получение значения переменной
	fmt.Print("введите строку: ")
	fmt.Fscanln(os.Stdin, &str)

	// вывод длины строки
	fmt.Println(str)
	fmt.Println(len(str))

}
