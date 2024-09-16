package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Написать функцию, которая принимает строку и возвращает ее длину.

	// получение значения переменной
	fmt.Print("введите строку: ")
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// вывод длины строки
	fmt.Println(len(text) - 2)

}
