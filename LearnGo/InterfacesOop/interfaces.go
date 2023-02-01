// Создан для поверхостного изучения интерфейсов.
package main

import (
	"fmt"
	"math"
)

type Square struct { // Структура квадрата
	sideLenght float64
}

type Circle struct { // Структура круга
	radius float64
}

type Shape interface { // Интерфейс с методом Area() для взаимодействия со структурами.
	Area() float64
}

func (s Square) Area() float64 { // Функция площади квадрата
	return s.sideLenght * s.sideLenght
}

func (c Circle) Area() float64 { // Функция площади круга
	return float64(math.Pi) * (float64(c.radius) * float64(c.radius))
}

func printShapeArea(s Shape) { // Функция вывода площадей.
	fmt.Printf("Площадь фигуры: %.2f см\n", s.Area())
}

func main() {
	square := Square{10.0} // Экземпляр структуры Square
	circle := Circle{5.5}

	printShapeArea(square)
	printShapeArea(circle)
}