package main

import (
	"fmt"
	"math"
	"errors"
)

type circle struct{ radius int }
type triangle struct{ sideOne, sideTwo, base int }
type rectangle struct{ sideOne, sideTwo int }

func main() { printFigure() }

func printFigure() {
var sideOne, sideTwo int
	for {
		var radius int
		fmt.Print("Введите радиус круга: ")
		fmt.Scan(&radius)
		if !validator(radius, 1, 1) { continue }
		сircle := circle{radius}
		сircle.printInfo()
		break
	} 
	for {
		var base int
		fmt.Print("Введите стороны и основание треугольника через пробел\nФормат - Сторона1 Сторона2 Основание: ")
		fmt.Scan(&sideOne, &sideTwo, &base)
		if !validator(sideOne, sideTwo, base) { continue }
		triangle := triangle{sideOne, sideTwo, base}
		triangle.printInfo()
		break
	} 
	for {
		fmt.Print("Введите ширину и длинну прямоугольника через пробел: ")
		fmt.Scan(&sideOne, &sideTwo)
		if !validator(sideOne, sideTwo, 1) { continue }
		rectangle := rectangle{sideOne, sideTwo}
		rectangle.printInfo()
		break
	}
}

func (c *circle) printInfo() {
	fmt.Printf("Площадь круга с радиусом %d см. равна %d см. в кв.\n\n", c.radius, int(float64(c.radius)*float64(c.radius)*math.Pi))
}

func (t *triangle) printInfo() {
	per := t.sideOne + t.sideTwo + t.base
	area := per / 2 * ((per/2 - t.sideOne) * (per/2 - t.sideTwo) * (per/2 - t.base))
	fmt.Printf("Площадь треугольника с периметром %d см. равна %d см. в кв.\n\n", per, area)
}

func (r *rectangle) printInfo() {
	per := r.sideOne*2 + r.sideTwo*2
	area := r.sideOne * r.sideTwo
	fmt.Printf("Площадь прямоугольника при периметре %d см. равна %d см. в кв.\n\n", per, area)
}

func validator(one, two, three int) bool {
	if one <= 0 || two <= 0 || three <= 0 {
		fmt.Println(errors.New("\nЗначение не может быть меньше или равно 0!\n"))
		return false
	}
	return true
}
