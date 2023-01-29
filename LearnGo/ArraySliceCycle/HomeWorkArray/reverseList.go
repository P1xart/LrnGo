// Обучение переворачиванию списков (Домашнее задание). 
// Не выполнено, на грани подсмотрено в гугле. Потрачено около 2 часов.
package main

import "fmt"

func main() {
	var list = []string{
		"Нуль", // 0
		"Один", // 1
		"Два",  // 2
		"Три",  // 3
	}
	list = reverseList(list)
	fmt.Println(list)
}
func reverseList(list []string) []string {
	last := len(list) - 1              // Переменная для последнего элемента
	for i := 0; i < len(list)/2; i++ { // Цикл работает, пока i меньше, чем длина списка уменьшенная в 2 раза.
		list[i], list[last-i] = list[last-i], list[i] // 0 = 3 -> 3 = 0; 1 = 2 -> 2 = 1...
	}
	return list
}
