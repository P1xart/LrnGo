// Домашняя работа. Практиктика мапов. АЛКОГОЛЬ - ЗЛО!
package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[0] = "Нуль"
	m[0] = "Ноль"
	m[1] = "Один"
	m[2] = "Два"
	m[3] = "Три"
	delete(m, 0)

	m2 := map[string]bool{
		"Виски":      true,
		"Пиво":       false,
		"Vodka":      false,
		"Шампанское": true,
		"Мартини":    true,
	}

	for index := range m2 {
		if m2[index] == true {
			fmt.Printf("%s - %s\n", index, "Пью")
		} else {
			fmt.Printf("%s - %s\n", index, "Не пью")
		}

	}

	fmt.Println()
	for index := range m {
		fmt.Printf("%d - %s\n", index, m[index])
	}
}
