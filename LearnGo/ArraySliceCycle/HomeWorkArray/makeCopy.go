// Обучение встроенным функция make() и copy(). Вспомнил append()
package main

import "fmt"

func main() {
	array := []string{"I", "Love", "Python"}
	array2 := make([]string, 2) // make([]type, длинна среза, длина массива на который указывает срез)
	copy(array2, array) // copy(dst array, src array)
	array2 = append(array2, "Go")
	fmt.Println("До:", array, "| После:", array2)
}