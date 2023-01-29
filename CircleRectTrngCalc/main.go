// CircleRectTrngCalc - КругПрямоугольникТреугольникКалькулятор
// Создан в целях изучения фундаментальных структур, методов, пакетов и возможностей языка Go в общем.
// Калькулятор высчитывает площадь и периметр круга, прямоугольника и треугольника.
package main

import ( // errors - валидаторы, fmt - ввод/вывод, math - Pi и квадратный корень, os - для выхода из программы в валидаторе
	"errors"
	"fmt"
	"math"
	"os"
)

var radius, height, width, base, side, side2 int // Обьявление всех переменных глобально в int

func main() {
	calcPrintCircleArea()
	calcPrintRectAreaPer() // Вызов всех функций
	calcPrintTrngAreaPer()
}
func calcPrintCircleArea() { // Вычисление и вывод площади круга по его радиусу
	fmt.Print("Введите радиус круга в сантиметрах: ")
	fmt.Scan(&radius)
	if radius <= 0 { // Валидатор круга
		fmt.Println(errors.New("Радиус не может быть меньше или равным нулю!"))
		os.Exit(0)
	}
	circleArea := float64(radius) * float64(radius) * math.Pi
	fmt.Printf("Площадь круга равна %.3f см. в кв.\n\n", circleArea)
}
func calcPrintRectAreaPer() { // Вычисление и вывод площади с периметром прямоугольника
	fmt.Print("Введите высоту и ширину прямоугольника в см. через пробел: ")
	fmt.Scan(&height, &width)
	if height <= 0 || width <= 0 { // Валидатор прямоугольника
		fmt.Println(errors.New("Высота или ширина не могут быть меньше или равны нулю!"))
		os.Exit(0)
	}
	rectArea := height * width
	rectPer := height*2 + width*2
	fmt.Printf("Площать прямоугольника равна %d см. в кв.\n", rectArea)
	fmt.Printf("Периметр прямоугольника равен %d см.\n\n", rectPer)
}
func calcPrintTrngAreaPer() { // Вычисление и вывод площади с периметром треугольника
	fmt.Print("Введите основание и cтороны треугольника (a, b, c) в см. через пробел: ")
	fmt.Scan(&base, &side, &side2)
	if base <= 0 || side <= 0 || side2 <= 0 { // Валидатор треугольника
		fmt.Println(errors.New("Основание или стороны не могут быть меньше или равны нулю!"))
		os.Exit(0)
	}
	semiPer := (base + side + side2) / 2
	trngArea := math.Sqrt(float64(semiPer) * (float64(semiPer) - float64(side)) * (float64(semiPer) - float64(base)) * (float64(semiPer) - float64(side2)))
	// Формула Герера для исчисления площади треугольника по всем 3 известным сторонам
	trngPer := base + side + side2
	fmt.Printf("Площать треугольника равна %.3f см. в кв.\n", trngArea)
	fmt.Printf("Периметр треугольника равен %d см.\n", trngPer)
}
