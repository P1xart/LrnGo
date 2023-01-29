// Домашнее задание. Изучение некоторых действий с массивами/срезами.
package main

import "fmt"

func main() {
	array := []int{0, 1, 2, 3, 4, 5}
	array = append(array[:1], array[1+1:]...) // Удаление. append(a[:i], a[i+1:]...)
	fmt.Println("Delete:", array)

	array = []int{0, 1, 2, 3, 4, 5}
	array2 := []int{6, 7, 8, 9, 10}
	array = append(array, array2...) // Распаковка. append(a, a2...). ... - оператор распаковки
	fmt.Println("Unpacking:", array)
}