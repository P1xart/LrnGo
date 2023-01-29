// Обучение передаче массива (возможно пустого), как аргумента и его получение с помощью return. Работа с исходно-пустым списком и добавление в него элементов.
package main

import "fmt"

func main() {
	var arr [3]int
	arr = fillArray(arr)
	fmt.Println(arr)
}
func fillArray(arr [3]int) [3]int {
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	return arr
}