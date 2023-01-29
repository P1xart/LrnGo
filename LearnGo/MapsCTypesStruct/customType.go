// Изучение кастомных типов.
package main

import "fmt"

type age int // Обьявление нового кастомного типа.

func main() {
	myAge := age(15) // age(15) дает понять что это наш кастомный тип age, а не int.

	fmt.Printf("Я совершеннолетний?\n%s", myAge.isAdult()) // Как можем видеть, myAge играет роль обьекта
} // а isAdult - роль метода обьекта.

func (a age) isAdult() string { // Метод c ресивером.
	if a >= 18 {
		return "ДА"
	} else {
		return "НЕТ"
	}
}
