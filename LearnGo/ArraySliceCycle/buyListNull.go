// Обучение работы с пустыми срезами и изменениям в них, переносом, добавлениям элементов.
package main

import "fmt"

func main() {
	var sliceList = [5]string{
		"Станцевать",
		"Покурить",
		"Похлопать",
		"Сбегать",
		"Отжиматься",
	}
	var listNull []string
	fmt.Println(listNull == nil)

	listNull = sliceList[:]
	fmt.Println(listNull == nil)
	changeList(listNull)
	for index := range sliceList {
		fmt.Printf("%d. %s\n", index+1, sliceList[index])
	}
}
func changeList(listNull []string) {
	listNull[3] = "Сходить"
	listNull[2] = "Апплодировать"
}