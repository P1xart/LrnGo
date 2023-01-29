// Игра "Угадай число". Создана для изучения цикла for, break, continue, библиотеки rand и повторения валидаторов.
package main

import (
	"errors" // Валидатор 0 < userNum < 100
	"fmt"
	"math/rand" // Rand.Intn(n) (Генерация числа [0,n))
	"time" // Для привязки rand.Intn к Unix времени системы и генерации рандомного числа.
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Генерация нового сида для постоянной генерации нового числа.
	num := rand.Intn(100)
	hearth := 4
	for { // Аналог while() в python.
		userNum := 0
		fmt.Print("Введите число: ")
		fmt.Scan(&userNum)
		if 0 > userNum || userNum > 100 { // Валидатор
			fmt.Println(errors.New("Число не может быть меньше 0 или больше 100!"))
			continue
		}
		if userNum == num { // Выйгрыш
			fmt.Println("\nWIN!")
			break
		} else if hearth == 0 { // Пройгрыш
			fmt.Println("\nGAME OVER!")
			break
		} else if num < userNum { // Меньше загаданного
			fmt.Printf("Число меньше загаданного! Осталось попыток: %d\n", hearth)
			hearth -= 1
			continue
		} else if num > userNum { // Больше загаданного
			fmt.Printf("Число больше загаданного! Осталось попыток: %d\n", hearth)
			hearth -= 1
			continue
		}
	}
}
