// Создан для изучения мапов. Аналог словарей в python.
package main

import "fmt"

func main() {
	cars := map[string]int {
		"BMW":10000,
		"Audi":8000,
		"Mercedes":7500,
		"Shevrolet":5000,
	}
	fmt.Println(cars["BMW"]) // Вывод: 10000

	delete(cars, "BMW") // Удаляем BMW
	fmt.Println(cars["BMW"]) // Вывод: 0, так как в 12 строке удалили BMW

	cars2 := cars // Копируем cars в новую переменную cars2

	delete(cars, "Shevrolet") // Удаляем шевролет
	fmt.Println(cars) 
	fmt.Println(cars2) // Оба значения индентичны, ибо указатели удалили значение как в cars, так и в cars2
}