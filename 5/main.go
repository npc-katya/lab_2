package main

import "fmt"

// создание структуры Rectangle
type Rectangle struct {
	x float64
	y float64
}

func main() {
	// Создать структуру Rectangle и реализовать метод для вычисления площади прямоугольника.

	rectangle := Rectangle{x: 3.5, y: 4}

	fmt.Printf("стороны прямоуголькика:  %f, %f\n", rectangle.x, rectangle.y)

	area(rectangle)

}

// функция для вычисления площади прямоугольника
func area(rectangle Rectangle) {
	x := rectangle.x * rectangle.y
	fmt.Println("площадь прямоугольника: ", x)
}
